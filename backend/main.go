package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func main() {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateApp(r)

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	port, found := os.LookupEnv("APP_PORT")
	if !found {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
