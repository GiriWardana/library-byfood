package controllers

import (
	"log"
	"net/http"

	"backend/models"
	"backend/services"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := services.DB.Find(&books).Error; err != nil {
		log.Printf("ERROR [GET /books]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	log.Printf("SUCCESS [GET /books]: Returned %d books", len(books))
	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book to create"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("ERROR [POST /books]: Bind JSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	// Validate required fields
	if book.Title == "" || book.Author == "" || book.Year == 0 {
		log.Printf("ERROR [POST /books]: Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title, Author, and Year are required"})
		return
	}

	if err := services.DB.Create(&book).Error; err != nil {
		log.Printf("ERROR [POST /books]: DB create failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	log.Printf("SUCCESS [POST /books]: Created book ID %d", book.Id)
	c.JSON(http.StatusOK, book)
}

// GetBook godoc
// @Summary Get a single book by ID
// @Description Get details of a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := services.DB.First(&book, id).Error; err != nil {
		log.Printf("ERROR [GET /books/%s]: Not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	log.Printf("SUCCESS [GET /books/%s]: Found book", id)
	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update book details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book data"
// @Success 200 {object} models.Book
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := services.DB.First(&book, id).Error; err != nil {
		log.Printf("ERROR [PUT /books/%s]: Not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var updateData models.Book
	if err := c.ShouldBindJSON(&updateData); err != nil {
		log.Printf("ERROR [PUT /books/%s]: Bind JSON error: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	// Validate required fields
	if updateData.Title == "" || updateData.Author == "" || updateData.Year == 0 {
		log.Printf("ERROR [PUT /books/%s]: Missing required fields", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title, Author, and Year are required"})
		return
	}

	book.Title = updateData.Title
	book.Author = updateData.Author
	book.Year = updateData.Year

	if err := services.DB.Save(&book).Error; err != nil {
		log.Printf("ERROR [PUT /books/%s]: Save failed: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	log.Printf("SUCCESS [PUT /books/%s]: Book updated", id)
	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := services.DB.First(&book, id).Error; err != nil {
		log.Printf("ERROR [DELETE /books/%s]: Not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := services.DB.Delete(&book).Error; err != nil {
		log.Printf("ERROR [DELETE /books/%s]: Delete failed: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	log.Printf("SUCCESS [DELETE /books/%s]: Book deleted", id)
	c.JSON(http.StatusOK, gin.H{"success": "Record deleted"})
}
