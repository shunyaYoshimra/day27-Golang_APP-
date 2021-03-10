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
	// seting environment
	flag.Parse()
	// create gin Engine
	r := gin.Default()

	app := new(apps.Application)
	// configure app followed environment
	if *env == "production" {
		app.CreateProductionApp(r)
	} else {
		app.CreateApp(r)
	}

	if err := godotenv.Load("../.env"); err != nil {
		// path for air hot reload
		if err := godotenv.Load("./.env"); err != nil {
			panic(err)
		}
	}

	// configure server port
	port, found := os.LookupEnv("APP_PORT")
	if !found {
		port = "8080"
	}

	// start engine !!
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
