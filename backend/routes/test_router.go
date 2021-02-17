package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/controllers"
)

func NewTestRouter(g *gin.RouterGroup) {
	// init controllers
	userController := controllers.NewUserController()
	sessionController := controllers.NewSessionController()
	profileController := controllers.NewProfileController()
	contactController := controllers.NewContactController()
	abroadController := controllers.NewAbroadController()
	occupationController := controllers.NewOccupationController()
	questionController := controllers.NewQuestionController()
	answerController := controllers.NewAnswerController()
	postController := controllers.NewPostController()
	favoriteController := controllers.NewFavoriteController()
	articleController := controllers.NewArticleController()
	checkController := controllers.NewCheckController()
	// routes
	{
		g.POST("/signup", userController.Signup)
		g.POST("/login", sessionController.Login)
		// routes for user controller
		g.GET("/users", userController.Index)
		g.GET("/user/:id", userController.Show)
		g.GET("/get_me", userController.GetMe)
		//  routes for profile controller
		g.GET("/profile/:id", profileController.Show)
		g.GET("/my_profile", profileController.GetMyProfile)
		g.POST("/profiles", profileController.Create)
		g.PUT("/profiles", profileController.Update)
		// routes for contact controller
		g.GET("/contact/:id", contactController.Show)
		g.GET("/my_contact", contactController.GetMyContact)
		g.POST("/contacts", contactController.Create)
		g.PUT("/contacts", contactController.Update)
		// routes for abroad controller
		g.GET("/abroads", abroadController.Index)
		g.GET("/abroad/:id", abroadController.Show)
		g.GET("/my_abroad", abroadController.MyAbroad)
		g.POST("/abroads", abroadController.Create)
		g.PUT("/abroads", abroadController.Update)
		// routes for occupation controller
		g.GET("/occupations", occupationController.Index)
		g.GET("/occupation/:id", occupationController.Show)
		g.GET("/my_occupation", occupationController.MyOccupation)
		g.POST("/occupations", occupationController.Create)
		g.PUT("/occupations", occupationController.Update)
		// routes for question controller
		g.GET("/questions", questionController.Index)
		g.GET("/question/:id", questionController.Show)
		g.GET("/questions/:id", questionController.UserQuestions)
		g.GET("/searched_questions/:about", questionController.SearchedIndex)
		g.POST("/questions", questionController.Create)
		g.PUT("/questions/:id", questionController.Update)
		g.PUT("/change_status/:id", questionController.ChangeStatus)
		g.DELETE("/questions/:id", questionController.Delete)
		// routes for answer controller
		g.GET("/answers/:id", answerController.Index)
		g.POST("/answers/:id", answerController.Create)
		g.PUT("/answers", answerController.Update)
		// routes for post controller
		g.GET("/posts", postController.Index)
		g.GET("/posts/:id", postController.UserPosts)
		g.GET("/searched_posts/:keyword", postController.SearchedIndex)
		g.GET("/post/:id", postController.Show)
		g.GET("/images/:id", postController.ImagesOfPost)
		g.POST("/posts/:tags", postController.Create)
		g.DELETE("/posts/:id", postController.Delete)
		// routes for favorite controller
		g.GET("/favorites", favoriteController.MyFavorites)
		g.GET("/favorites/:id", favoriteController.FavoritesOfUser)
		g.POST("/favorites/:id", favoriteController.Create)
		g.DELETE("/favorites/:id", favoriteController.Delete)
		// routes for article controller
		g.GET("/articles", articleController.Index)
		g.GET("/articles/:id", articleController.UserArticles)
		g.GET("/article/:id", articleController.Show)
		g.POST("/articles", articleController.Create)
		g.PUT("/articles/:id", articleController.Update)
		g.DELETE("/articles/:id", articleController.Delete)
		// routes for check controller
		g.GET("/checks/:id", checkController.ChecksOfArticle)
		g.GET("/user_checks/:id", checkController.ChecksOfUser)
		g.POST("/checks/:id", checkController.Create)
		g.DELETE("/checks/:id", checkController.Delete)
	}
}
