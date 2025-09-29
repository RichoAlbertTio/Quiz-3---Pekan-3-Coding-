package main

import (
	"log"
	"os"

	"quiz-13/internal/auth"
	"quiz-13/internal/books"
	"quiz-13/internal/categories"
	"quiz-13/internal/database"
	"quiz-13/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Inisialisasi database
	database.ConnectDB()

	r := gin.Default()

	// Users
	r.POST("/api/users/login", users.Login)

	// Categories
	cat := r.Group("/api/categories")
	cat.Use(auth.JWTMiddleware())
	{
		cat.GET("/", categories.GetCategories)
		cat.POST("/", categories.CreateCategory)
	}

	// Books
	book := r.Group("/api/books")
	book.Use(auth.JWTMiddleware())
	{
		book.GET("/", books.GetBooks)
		book.POST("/", books.CreateBook)
	}

	// Dapatkan port dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	r.Run(":" + port)
}
