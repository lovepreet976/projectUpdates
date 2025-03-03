package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterOwnerNew(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the role is "owner"
	if input.Role != "owner" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role, must be 'owner'"})
		return
	}

	// Hash password before saving (assuming HashPassword function exists)
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}
	input.Password = hashedPassword

	// Create the new owner
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create owner"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New owner registered successfully", "owner": input})
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
