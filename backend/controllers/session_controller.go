package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
)

type SessionController struct {
	Repository repositories.UserRepository
}

func NewSessionController() SessionController {
	return SessionController{
		Repository: repositories.NewUserRepository(),
	}
}

func (sc *SessionController) Login(c *gin.Context) {
	email := c.PostForm("email")
	formPassword := c.PostForm("password")
	if email == "" || formPassword == "" {
		res := response.BadRequest("全ての情報を入力して下さい(All information should be filled in)")
		c.JSON(res.Status, res)
	} else {
		if user, err := sc.Repository.FindByEmail(email); err != nil {
			res := response.NotFound("このEmailは登録されていません(This e-mail hasn't been registated yet.)")
			c.JSON(res.Status, res)
		} else {
			dbPassword := user.Password
			if err := middleware.CompareHashAndPassword(dbPassword, formPassword); err != nil {
				res := response.BadRequest("パスワードが違います(Password is incorrect)")
				c.JSON(res.Status, res)
			} else {
				id := user.ID
				middleware.Login(c, id)
				res := response.SuccessResponse("")
				c.JSON(res.Status, res)
			}
		}
	}
}

func (sc *SessionController) Logout(c *gin.Context) {
	middleware.Logout(c)
	c.Redirect(302, "/app/#/login")
}

func (sc *SessionController) Check(c *gin.Context) {
	middleware.SessionBool(c)
}
