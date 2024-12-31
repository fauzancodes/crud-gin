package repository

//Directory: /app/repository/product.go
//This file is used to perform ORM queries to the database

import (
	"fmt"

	"github.com/fauzancodes/crud-gin/app/config"
	"github.com/fauzancodes/crud-gin/app/models"
	"github.com/google/uuid"
)

func CreateProduct(data models.CRUDProduct) (models.CRUDProduct, error) {
	//This ORM function is used to insert new data into the database.
	//The database query formed from the following ORM functions is as follows:
	//INSERT INTO crud_products (code, name, description, status, price) VALUES (?, ?, ?, ?, ?);
	//For the id, created_at, updated_at columns, they will be filled in automatically because we have set the default values ​​in the models, otherwise these columns will have a NULL value, while for the deleted_at column, the value will be NULL
	err := config.DB.Create(&data).Error

	return data, err
}

func GetProducts(search string) (responses []models.CRUDProduct, err error) {
	//This ORM function is used to get all data from the database
	var where string
	var values []any
	//The following condition is used to add the WHERE parameter to the database query that will be formed to filter the data that will be retrieved based on the search entered by the client request
	if search != "" {
		//The "?" sign is used to perform parameterized queries to prevent sql injection. Please search more on google about parameterized queries and sql injection if you don't know
		where += "code ILIKE ? OR name ILIKE ? OR description ILIKE ?"
		values = append(values, fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	//The database query formed from the following ORM functions is as follows:
	// SELECT * FROM crud_products WHERE code ILIKE ? OR name ILIKE ? OR description ILIKE ?
	err = config.DB.Where(where, values...).Find(&responses).Error

	return
}

func GetProductByID(id uuid.UUID) (response models.CRUDProduct, err error) {
	//This ORM function is used to get a data from the database based on id
	//The database query formed from the following ORM functions is as follows:
	//SELECT * FROM crud_products WHERE id = ?
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func UpdateProduct(data models.CRUDProduct) (models.CRUDProduct, error) {
	//This ORM function is used to update a data from the database based on id
	//The database query formed from the following ORM functions is as follows:
	//UPDATE crud_products SET code = ?, name = ?, description = ?, status = ?, price = ? WHERE id = ?
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteProduct(data models.CRUDProduct) error {
	//This ORM function is used to delete a data from the database based on id
	//The database query formed from the following ORM functions is as follows:
	//DELETE FROM crud_products WHERE id = ?
	err := config.DB.Delete(&data).Error

	return err
}
