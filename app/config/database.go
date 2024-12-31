package config

//Directory: /app/config/database.go
//This file is used to set up the database

import (
	"fmt"
	"log"

	"github.com/fauzancodes/crud-gin/app/models"
	"gorm.io/driver/postgres" //This package is used because we use PostgreSQL as our database
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() *gorm.DB {
	host := LoadConfig().DatabaseHost         //Get database host from environment variable
	user := LoadConfig().DatabaseUsername     //Get database username from environment variable
	password := LoadConfig().DatabasePassword //Get database password from environment variable
	name := LoadConfig().DatabaseName         //Get database name from environment variable
	port := LoadConfig().DatabasePort         //Get database port from environment variable

	//Connect to the database
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, name, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if LoadConfig().EnableDatabaseAutomigration {
		//Run database automigration asynchronously
		go RunAutoMigration()
	}

	log.Printf("Connected to Database: %v", name)

	return DB
}

func RunAutoMigration() {
	err := DB.AutoMigrate(
		//List of models that will be structured into the database automatically by GORM when automigration runs
		&models.CRUDProduct{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
