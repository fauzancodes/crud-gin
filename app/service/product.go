package service

//Directory: /app/service/product.go
//This file is used to perform the business logic required for data processing either from the client to the database or vice versa from the database to the client

import (
	"errors"
	"net/http"

	"github.com/fauzancodes/crud-gin/app/dto"
	"github.com/fauzancodes/crud-gin/app/models"
	"github.com/fauzancodes/crud-gin/app/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateProduct(request dto.ProductRequest) (response models.CRUDProduct, statusCode int, err error) {
	//This function moves data from the client request to the models which will be sent to the database via the ORM
	data := models.CRUDProduct{
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
		Status:      request.Status,
		Price:       request.Price,
	}

	//Sending client request data to ORM
	response, err = repository.CreateProduct(data)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	//If the process is successful then the function is complete
	statusCode = http.StatusCreated
	return
}

func GetProductByID(id string) (response models.CRUDProduct, statusCode int, err error) {
	//Convert uuid from string format to uuid format
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	//Retrieving data from database using ORM
	response, err = repository.GetProductByID(parsedUUID)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	//If the process is successful then the function is complete
	statusCode = http.StatusOK
	return
}

func GetProducts(search string) (responses []models.CRUDProduct, statusCode int, err error) {
	//Retrieving data from database using ORM
	responses, err = repository.GetProducts(search)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	//If the process is successful then the function is complete
	statusCode = http.StatusOK
	return
}

func UpdateProduct(id string, request dto.ProductRequest) (response models.CRUDProduct, statusCode int, err error) {
	//Convert uuid from string format to uuid format
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	//Retrieving data from database using ORM
	data, err := repository.GetProductByID(parsedUUID)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	//Check whether the "Code" field is filled in by the client or whether this field is sent or not by the frontend, if not then the data from the database will not be replaced, if filled in then the data from the database will be replaced with the contents from the client
	if request.Code != "" {
		data.Code = request.Code
	}
	//Check whether the "Name" field is filled in by the client or whether this field is sent or not by the frontend, if not then the data from the database will not be replaced, if filled in then the data from the database will be replaced with the contents from the client
	if request.Name != "" {
		data.Name = request.Name
	}
	//Check whether the "Description" field is filled in by the client or whether this field is sent or not by the frontend, if not then the data from the database will not be replaced, if filled in then the data from the database will be replaced with the contents from the client
	if request.Description != "" {
		data.Description = request.Description
	}

	//Because the "Price" can be 0, this field cannot be conditioned to be skipped like the other fields above, therefore, the data from the database will be directly replaced with the contents of the client. Therefore, the frontend must send this field, otherwise it will accidentally delete the value in this field
	data.Price = request.Price

	//Because the "Status" only has 2 values, true or false, this field cannot be conditioned to be skipped like the other fields above, therefore, the data from the database will be directly replaced with the contents of the client. Therefore, the frontend must send this field, otherwise it will accidentally change the value in this field
	data.Status = request.Status

	//Sending updated data to ORM
	response, err = repository.UpdateProduct(data)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to update data: " + err.Error())
		statusCode = http.StatusInternalServerError
	}

	//If the process is successful then the function is complete
	statusCode = http.StatusOK
	return
}

func DeleteProduct(id string) (statusCode int, err error) {
	//Convert uuid from string format to uuid format
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	//Retrieving data from database using ORM
	data, err := repository.GetProductByID(parsedUUID)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	//Deleting data via ORM
	err = repository.DeleteProduct(data)
	if err != nil {
		//If the process fails, it will return an error
		err = errors.New("failed to delete data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	//If the process is successful then the function is complete
	statusCode = http.StatusOK
	return
}
