package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/controllers"
	"github.com/shunyaYoshimra/day27/backend/middleware"
)

func NewRouter(g *gin.RouterGroup) {
	userController := controllers.NewUserController()
	sessionController := controllers.NewSessionController()
	profileController := controllers.NewProfileController()
	contactController := controllers.NewContactController()
	{
		g.POST("/signup", userController.Signup)
		g.POST("/login", sessionController.Login)
		g.GET("/logout", sessionController.Logout)
		g.GET("/session_check", sessionController.Check)
	}
	// routes only for login user
	l := g.Group("/", middleware.SessionCheck())
	{
		// routes for user controller
		l.GET("/users", userController.Index)
		l.GET("/user/:id", userController.Show)
		l.GET("/get_me", userController.GetMe)
		//  routes for profile controller
		l.GET("/profile/:id", profileController.Show)
		l.GET("/my_profile", profileController.GetMyProfile)
		l.POST("/profiles", profileController.Create)
		l.PUT("/profiles", profileController.UpdateDescription)
		// routes for contact controller
		l.GET("/contact/:id", contactController.Show)
		l.GET("/my_contact", contactController.GetMyContact)
		l.POST("/contacts", contactController.Create)
		l.PUT("/contacts", contactController.Update)
	}
}
