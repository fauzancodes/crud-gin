package dto

//Directory: /app/dto/product.go
//This file is used to store DTO (data transfer object) which is a replica structure of the client's http request body

import validation "github.com/go-ozzo/ozzo-validation"

type ProductRequest struct {
	//This struct will be used to bind http requests from clients.
	//The "json:" tag is the name of the field that will be matched against the json field in the http request body from the client
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
	Price       float64 `json:"price"`
}

func (request ProductRequest) Validate() error {
	//This is a function to validate requests from clients
	return validation.ValidateStruct(
		&request,
		//This validation indicates that the "Name" field is required and cannot be empty
		validation.Field(&request.Name, validation.Required),
		//This validation indicates that the "Price" field must not be less than 0.0
		validation.Field(&request.Price, validation.Min(0.0)),
		//There are many other validations that can be done with this third party package, please explore it yourself
	)
}
