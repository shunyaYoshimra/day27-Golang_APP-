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
	abroadController := controllers.NewAbroadController()
	occupationController := controllers.NewOccupationController()
	questionController := controllers.NewQuestionController()
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
		l.PUT("/profiles", profileController.Update)
		// routes for contact controller
		l.GET("/contact/:id", contactController.Show)
		l.GET("/my_contact", contactController.GetMyContact)
		l.POST("/contacts", contactController.Create)
		l.PUT("/contacts", contactController.Update)
		// routes for abroad controller
		l.GET("/abroads", abroadController.Index)
		l.GET("/abroad/:id", abroadController.Show)
		l.GET("/my_abroad", abroadController.MyAbroad)
		l.POST("/abroads", abroadController.Create)
		l.PUT("/abroads", abroadController.Update)
		// routes for occupation controller
		l.GET("/occupations", occupationController.Index)
		l.GET("/occupation/:id", occupationController.Show)
		l.GET("/my_occupation", occupationController.MyOccupation)
		l.POST("/occupations", occupationController.Create)
		l.PUT("/occupations", occupationController.Update)
		// routes for question controller
		l.GET("/questions", questionController.Index)
		l.GET("/question/:id", questionController.Show)
		l.GET("/questions/:id", questionController.UserQuestions)
		l.GET("/searched_questions", questionController.SearchedIndex)
		l.POST("/questions", questionController.Create)
		l.PUT("/questions/:id", questionController.Update)
		l.PUT("/chage_status", questionController.ChangeStatus)
		l.DELETE("/questions/:id", questionController.Delete)
	}
}
