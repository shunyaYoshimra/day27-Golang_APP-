package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func Login(c *gin.Context, id int) {
	session := sessions.Default(c)
	session.Set("userID", id)
	session.Save()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func SessionBool(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("userID")
	if sessionID != nil {
		c.JSON(200, true)
	} else {
		c.JSON(200, false)
	}
}

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("userID")
		if sessionID == nil {
			c.JSON(200, gin.H{
				"message": "ログインしていません",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetSession(c *gin.Context) (id int) {
	session := sessions.Default(c)
	sessionID := session.Get("userID")
	id, _ = sessionID.(int)
	return
}
