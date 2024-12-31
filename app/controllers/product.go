package controllers

//Directory: /app/controllers/product.go
//This file is used to handle http requests from clients and responses to clients

import (
	"net/http"

	"github.com/fauzancodes/crud-gin/app/dto"
	"github.com/fauzancodes/crud-gin/app/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	//Bind http request body from client to dto struct
	var request dto.ProductRequest
	if err := c.Bind(&request); err != nil {
		//If the process fails, it will return an error
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)

		return
	}

	//Validating requests from clients
	if err := request.Validate(); err != nil {
		//If the process fails, it will return an error
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)

		return
	}

	//Sending requests from clients for further processing
	result, statusCode, err := service.CreateProduct(request)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to create product",
				Error:   err.Error(),
			},
		)

		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to create product",
			Data:    result,
		},
	)
}

func GetProducts(c *gin.Context) {
	//Retrieving query params from client http request
	search := c.Query("search")

	//Sending requests from clients for further processing
	result, statusCode, err := service.GetProducts(search)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to create product",
				Error:   err.Error(),
			},
		)

		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to create product",
			Data:    result,
		},
	)
}

func GetProductByID(c *gin.Context) {
	//Retrieving path params from client http request
	id := c.Param("id")

	//Sending requests from clients for further processing
	result, statusCode, err := service.GetProductByID(id)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to create product",
				Error:   err.Error(),
			},
		)

		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to create product",
			Data:    result,
		},
	)
}

func UpdateProduct(c *gin.Context) {
	//Bind http request body from client to dto struct
	var request dto.ProductRequest
	if err := c.Bind(&request); err != nil {
		//If the process fails, it will return an error
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)

		return
	}

	//Retrieving path params from client http request
	id := c.Param("id")

	//Sending requests from clients for further processing
	data, statusCode, err := service.UpdateProduct(id, request)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to update data",
				Error:   err.Error(),
			},
		)

		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(
		http.StatusOK,
		dto.Response{
			Status:  200,
			Message: "Success to update data",
			Data:    data,
		},
	)
}

func DeleteProduct(c *gin.Context) {
	//Retrieving path params from client http request
	id := c.Param("id")

	//Sending requests from clients for further processing
	statusCode, err := service.DeleteProduct(id)
	if err != nil {
		//If the process fails, it will return an error
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to delete data",
				Error:   err.Error(),
			},
		)

		return
	}

	//If the process is successful, the function is complete and will return response data to the client
	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to delete data",
		},
	)
}
