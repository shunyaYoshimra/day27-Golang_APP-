package apps

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/migration"
	"github.com/shunyaYoshimra/day27/backend/routes"
)

type Application struct{}

func (a Application) CreateApp(r *gin.Engine) {
	configureAppDB()
	configureAPIEndpoint(r)
	configureView(r)
}

func (a Application) CreateTest(r *gin.Engine) {
	configureTestDB()
	configureAPIEndpoint(r)
}

func configureAppDB() {
	database.AppConnection()
	conn := database.GetDB()
	migration.AutoMigration(conn)
}

func configureTestDB() {
	database.TestConnection()
	conn := database.GetDB()
	migration.AutoMigration(conn)
}

func configureAPIEndpoint(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	g := r.Group("/api/")
	routes.NewRouter(g)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "The page you are looking for dosen't exist",
		})
	})
}

func configureView(r *gin.Engine) {
	r.Static("/src", "../frontend/dist")
	r.StaticFS("/app", http.Dir("../frontend/static"))
}
