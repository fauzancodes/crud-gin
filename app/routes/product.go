package routes

//Directory: /app/routes/product.go
//This file is used to handle all the required routing

import (
	"github.com/fauzancodes/crud-gin/app/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoute(api *gin.RouterGroup) {
	product := api.Group("/product")
	{
		//POST is used to add new data
		product.POST("", controllers.CreateProduct)
		//GET is used to retrieve data that has been stored
		product.GET("", controllers.GetProducts)
		//GET /:id is used to retrieve data that has been stored based on id
		product.GET("/:id", controllers.GetProductByID)
		//PATCH /:id is used to change data that has been stored based on id
		product.PATCH("/:id", controllers.UpdateProduct)
		//DELETE /:id is used to change data that has been stored based on id
		product.DELETE("/:id", controllers.DeleteProduct)
	}
}
