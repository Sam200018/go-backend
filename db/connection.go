package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost port=5432 user=samuelbautista password= dbname=billingdb sslmode=disable"
var DB *gorm.DB

func DbConnection() {
	var error error

	DB, error = gorm.Open(
		postgres.Open(DSN),
		&gorm.Config{},
	)

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
