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
		product.POST("", controllers.CreateProduct)       //POST is used to add new data
		product.GET("", controllers.GetProducts)          //GET is used to retrieve data that has been stored
		product.GET("/:id", controllers.GetProductByID)   //GET /:id is used to retrieve data that has been stored based on id
		product.PATCH("/:id", controllers.UpdateProduct)  //PATCH /:id is used to change data that has been stored based on id
		product.DELETE("/:id", controllers.DeleteProduct) //DELETE /:id is used to change data that has been stored based on id
	}
}
