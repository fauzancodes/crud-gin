package utils

//Directory: /app/pkg/utils/utils.go
//This file is used to store frequently used common helper functions

import "github.com/fauzancodes/crud-gin/app/dto"

func MakeHTTPError(messange string, statusCode int, err error) (response dto.Response) {
	//This function is used to create an http response error
	response = dto.Response{
		Status:  statusCode,
		Message: messange,
		Error:   err.Error(),
	}

	return
}

func MakeHTTPResponse(messange string, statusCode int, data any) (response dto.Response) {
	//This function is used to create an http response
	response = dto.Response{
		Status:  statusCode,
		Message: messange,
		Data:    data,
	}

	return
}
