package main

import (
	"github.com/gin-gonic/gin"
	"quiz-13/internal/auth"
	"quiz-13/internal/books"
	"quiz-13/internal/categories"
	"quiz-13/internal/database"
	"quiz-13/internal/users"
)

func main() {
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

	r.Run(":8080")
}
