package categories

import (
	"net/http"
	"time"

	"quiz-13/internal/database"
	"quiz-13/internal/models"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{} // sementara pakai memory, bisa diganti dengan database

func GetCategories(c *gin.Context) {
	username, _ := c.Get("username")
	_ = username // untuk sementara tidak digunakan
	
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func CreateCategory(c *gin.Context) {
	username, _ := c.Get("username")
	usernameStr, _ := username.(string)
	
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	category := models.Category{
		ID:         len(categories) + 1,
		Name:       input.Name,
		CreatedAt:  time.Now(),
		CreatedBy:  usernameStr,
		ModifiedAt: time.Now(),
		ModifiedBy: usernameStr,
	}
	
	categories = append(categories, category)
	
	// TODO: Simpan ke database
	_ = database.DB // placeholder untuk menggunakan database
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    category,
	})
}
