package controllers

import (
	"backend/models"
	"backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	_ = db.AutoMigrate(&models.Book{})
	services.SetTestDB(db)

	r := gin.Default()
	// Register routes
	r.GET("/books", GetBooks)
	r.POST("/books", CreateBook)
	r.GET("/books/:id", GetBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)

	return r
}
