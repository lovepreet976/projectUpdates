package controllers

import (
	"fmt"
	"library-management/config"
	"library-management/models"
	"library-management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var input struct {
		Name       string `json:"name" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
		Contact    string `json:"contact"`
		LibraryIDs []uint `json:"library_ids" binding:"required"` // List of library IDs
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password before storing
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}

	// Enforce role as "user" to prevent constraint violations
	role := "user"

	// Create user
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Contact:  input.Contact,
		Role:     role,
	}

	// Save user and check for errors
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user", "details": err.Error()})
		return
	}

	// Link user with selected libraries
	for _, libID := range input.LibraryIDs {
		var library models.Library
		if err := config.DB.First(&library, libID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Library ID %d not found", libID)})
			return
		}

		userLibrary := models.UserLibrary{
			UserID:    user.ID,
			LibraryID: libID,
		}

		if err := config.DB.Create(&userLibrary).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to associate user with library",
				"details": err.Error(),
			})
			return
		}
	}

	// Fetch linked libraries manually since GORM doesn't auto-load many-to-many relations
	var userLibraries []models.UserLibrary
	config.DB.Where("user_id = ?", user.ID).Find(&userLibraries)

	var linkedLibraries []models.Library
	for _, ul := range userLibraries {
		var library models.Library
		config.DB.First(&library, ul.LibraryID)
		linkedLibraries = append(linkedLibraries, library)
	}

	// Return user with linked libraries
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"ID":        user.ID,
			"Name":      user.Name,
			"Email":     user.Email,
			"Role":      user.Role,
			"Contact":   user.Contact,
			"Libraries": linkedLibraries,
		},
	})
}
