package books

import (
	"net/http"
	"time"

	"quiz-13/internal/database"
	"quiz-13/internal/models"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{}

func GetBooks(c *gin.Context) {
	username, _ := c.Get("username")
	_ = username // untuk sementara tidak digunakan
	
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func CreateBook(c *gin.Context) {
	username, _ := c.Get("username")
	usernameStr, _ := username.(string)
	
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year" binding:"required"`
		Price       int    `json:"price"`
		TotalPage   int    `json:"total_page" binding:"required"`
		CategoryID  int    `json:"category_id"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validasi release_year
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2025 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2025"})
		return
	}

	// Konversi thickness berdasarkan total_page
	var thickness string
	if input.TotalPage > 100 {
		thickness = "tebal"
	} else {
		thickness = "tipis"
	}

	book := models.Book{
		ID:          len(books) + 1,
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
		CategoryID:  input.CategoryID,
		CreatedAt:   time.Now(),
		CreatedBy:   usernameStr,
		ModifiedAt:  time.Now(),
		ModifiedBy:  usernameStr,
	}
	
	books = append(books, book)
	
	// TODO: Simpan ke database
	_ = database.DB // placeholder untuk menggunakan database
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"data":    book,
	})
}
