package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/controllers"
	"github.com/shunyaYoshimra/day27/backend/middleware"
)

func NewRouter(g *gin.RouterGroup) {
	userController := controllers.NewUserController()
	sessionController := controllers.NewSessionController()
	{
		g.POST("/signup", userController.Signup)
		g.POST("/login", sessionController.Login)
		g.GET("/logout", sessionController.Logout)
		g.GET("/session_check", sessionController.Check)
	}
	// routes only for login user
	l := g.Group("/", middleware.SessionCheck())
	{
		l.GET("/users", userController.Index)
		l.GET("/user/:id", userController.Show)
	}
}
