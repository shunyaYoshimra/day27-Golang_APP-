package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

var env = flag.String("env", "development", "")

func main() {
	flag.Parse()

	r := gin.Default()
	app := new(apps.Application)
	if *env == "production" {
		app.CreateProductionApp(r)
	} else {
		app.CreateApp(r)
	}

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
