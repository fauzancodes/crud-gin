package models

//Directory: /app/models/product.go
//This file is used to store models which are replica structures of tables in the database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CRUDProduct struct {
	//This struct is a model that will be structured into the database automatically by GORM when automigration runs.
	//The structure is created based on the "gorm:" tag, "default" is the default value, "type" is the data type, "column" is the column name.
	//The "json:" tag is the field name that will be used when the data is returned as a json response to the client
	ID          uuid.UUID       `gorm:"type:uuid;column:id;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt   time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
	Code        string          `json:"code" gorm:"type:varchar(50);column:code"`
	Name        string          `json:"name" gorm:"type:varchar(255);column:name"`
	Description string          `json:"description" gorm:"type:text;column:description"`
	Status      bool            `json:"status" gorm:"type:bool;column:status"`
	Price       float64         `json:"price" gorm:"type:float8;column:price"`
}

func (CRUDProduct) TableName() string {
	//This is the table name of this model that will be created by GORM when the automigration process runs
	return "crud_products"
}

//As an additional note, if you specify the value "column" in the "gorm:" tag and also do not specify the TableName() function, then GORM will create column names and table names for this model automatically with the following rules:
//- names will be snake_case for PascalCase models
//- names will be plural for English
