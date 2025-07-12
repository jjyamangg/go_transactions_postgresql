package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Parameters to generate the connection to the database in PostgreSQL
var DSN = "host=localhost user=postgres password=**postgreSQL.2025 dbname=golang port=5432"

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
		log.Println("Unable to establish a connection to the database!")
	} else {
		log.Println("Connection to the database!")
	}
}
