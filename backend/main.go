package main

import (
	"backend/controllers"
	"backend/services"

	_ "backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Allow requests from your frontend
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	services.ConnectDatabase()

	// Define your routes
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/clean-url", controllers.CleanURLHandler)

	r.Run(":8080")
}
