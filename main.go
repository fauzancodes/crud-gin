package main

//Directory: /main.go
//This file is the application entry file

import (
	"log"

	"github.com/fauzancodes/crud-gin/app/config"
	"github.com/fauzancodes/crud-gin/app/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" //This package is used to load environment variables automatically every time the application runs
)

func main() {
	//App initiation
	app := Init()

	//Get port from environment variable
	port := config.LoadConfig().Port

	log.Println("Server running on port:" + port)

	//Run the application on the specified port
	app.Run(":" + port)
}

func Init() *gin.Engine {
	//Gin initiation
	app := gin.Default()

	//Database initiation
	config.Database()

	//Routing initiation
	routes.Route(app)

	return app
}
