package infrastructure

import (
	"golang-crud-rest-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitializeDatabase(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&entities.Product{})
	log.Println("Connected to DB")
}
