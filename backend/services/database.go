// In services/database.go
package services

import (
	"backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetTestDB(db *gorm.DB) {
	DB = db
}

func ConnectDatabase() {
	dsn := "host=db user=user password=password dbname=library_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&models.Book{})
	database.AutoMigrate(&models.CleanURL{})

	log.Println("Database connection established successfully")

	DB = database
}
