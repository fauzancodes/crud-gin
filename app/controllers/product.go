package controllers

//Directory: /app/controllers/product.go
//This file is used to handle http requests from clients and responses to clients

import (
	"net/http"

	"github.com/fauzancodes/crud-gin/app/dto"
	"github.com/fauzancodes/crud-gin/app/pkg/utils"
	"github.com/fauzancodes/crud-gin/app/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	//Bind http request body from client to dto struct
	var request dto.ProductRequest
	if err := c.Bind(&request); err != nil {
		//If the process fails, it will return an error
		c.JSON(http.StatusUnprocessableEntity, utils.MakeHTTPError("Invalid request body", http.StatusUnprocessableEntity, err))
		return
	}

	//Validating requests from clients
	if err := request.Validate(); err != nil {
		//If the process fails, it will return an error
		c.JSON(http.StatusBadRequest, utils.MakeHTTPError("Invalid request value", http.StatusBadRequest, err))
		return
	}

	//Sending requests from clients for further processing
	result, statusCode, err := service.CreateProduct(request)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(statusCode, utils.MakeHTTPError("Invalid request value", statusCode, err))
		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(statusCode, utils.MakeHTTPResponse("Success to create product", statusCode, result))
}

func GetProducts(c *gin.Context) {
	//Retrieving query params from client http request
	search := c.Query("search")

	//Sending requests from clients for further processing
	result, statusCode, err := service.GetProducts(search)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(statusCode, utils.MakeHTTPError("Failed to get product", statusCode, err))
		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(statusCode, utils.MakeHTTPResponse("Success to get product", statusCode, result))
}

func GetProductByID(c *gin.Context) {
	//Retrieving path params from client http request
	id := c.Param("id")

	//Sending requests from clients for further processing
	result, statusCode, err := service.GetProductByID(id)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(statusCode, utils.MakeHTTPError("Failed to get product", statusCode, err))
		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(statusCode, utils.MakeHTTPResponse("Success to get product", statusCode, result))
}

func UpdateProduct(c *gin.Context) {
	//Bind http request body from client to dto struct
	var request dto.ProductRequest
	if err := c.Bind(&request); err != nil {
		//If the process fails, it will return an error
		c.JSON(http.StatusUnprocessableEntity, utils.MakeHTTPError("Invalid request body", http.StatusUnprocessableEntity, err))
		return
	}

	//Retrieving path params from client http request
	id := c.Param("id")

	//Sending requests from clients for further processing
	data, statusCode, err := service.UpdateProduct(id, request)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(statusCode, utils.MakeHTTPError("Failed to update data", statusCode, err))
		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(statusCode, utils.MakeHTTPResponse("Success to update data", statusCode, data))
}

func DeleteProduct(c *gin.Context) {
	//Retrieving path params from client http request
	id := c.Param("id")

	//Sending requests from clients for further processing
	statusCode, err := service.DeleteProduct(id)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(statusCode, utils.MakeHTTPError("Failed to delete data", statusCode, err))
		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(statusCode, utils.MakeHTTPResponse("Success to delete data", statusCode, nil))
}
