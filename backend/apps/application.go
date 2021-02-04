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
			"status": 404,
		})
	})
}

func configureView(r *gin.Engine) {
	r.Static("/dist", "../frontend/dist")
	r.StaticFS("/app", http.Dir("../frontend/static"))
}
