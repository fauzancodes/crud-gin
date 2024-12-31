package config

//Directory: /app/config/config.go
//This file is to set up all your environment variables, this is done to make it easier to manage your environment variables because they are grouped in one file, and when, for example, your environment variable changes name, then you don't need to change the name of your environment variable in all your coding files, just in this file

import (
	"os"
	"strconv"
)

type Config struct {
	Port                        string
	DatabaseUsername            string
	DatabasePassword            string
	DatabaseHost                string
	DatabasePort                string
	DatabaseName                string
	EnableDatabaseAutomigration bool
	//Automigration is a process for migrating (creating a structure) a database automatically, directly handled by GORM
}

func LoadConfig() (config *Config) {
	port := os.Getenv("PORT")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))

	return &Config{
		DatabaseUsername:            databaseUsername,
		DatabasePassword:            databasePassword,
		DatabaseHost:                databaseHost,
		DatabasePort:                databasePort,
		DatabaseName:                databaseName,
		EnableDatabaseAutomigration: enableDatabaseAutomigration,
		Port:                        port,
	}
}
