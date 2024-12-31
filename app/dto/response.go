package dto

//Directory: /app/dto/response.go
//This file is used to display the json response to the client uniformly.

type Response struct {
	//The "json:" tag is the field name that will be used when the data is returned as a json response to the client.
	//The "omitempty" option indicates that, if the field is empty, the field will not be displayed in the json response sent to the client
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
