package routes

//Directory: /app/routes/index.go
//This file is used to group all existing routing, and is also used to create routing versions

import "github.com/gin-gonic/gin"

func Route(app *gin.Engine) {
	api := app.Group("/v1")

	//List all available routing
	ProductRoute(api)
}
