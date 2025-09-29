package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"quiz-13/internal/auth"
	"quiz-13/internal/database"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// TODO: Cek database user untuk validasi yang sebenarnya
	// Sementara hardcode untuk testing
	if input.Username != "admin" || input.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := auth.GenerateToken(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// TODO: Update last login di database
	_ = database.DB // placeholder untuk menggunakan database

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    gin.H{
			"username": input.Username,
		},
	})
}
