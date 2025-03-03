package controllers

import (
	"fmt"
	"library-management/config"
	"library-management/models"
	"library-management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdmin(c *gin.Context) {
	var input struct {
		Name       string `json:"name" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
		Contact    string `json:"contact"`
		LibraryIDs []uint `json:"library_ids" binding:"required"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure only an existing owner can create an admin
	creatorID := c.GetUint("userID") // Extract userID from JWT
	var creator models.User

	if err := config.DB.First(&creator, creatorID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid owner, user not found"})
		return
	}

	// Ensure the creator is actually an owner
	if creator.Role != "owner" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only an owner can create an admin"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create new Admin user
	admin := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Contact:  input.Contact,
		Role:     "admin",
	}

	// Save admin first
	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create admin"})
		return
	}

	// Link admin with selected libraries
	for _, libID := range input.LibraryIDs {
		var library models.Library
		if err := config.DB.First(&library, libID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Library ID %d not found", libID)})
			return
		}

		// Add admin to library
		adminLibrary := models.UserLibrary{
			UserID:    admin.ID,
			LibraryID: libID,
		}
		config.DB.Create(&adminLibrary)
	}

	// ✅ Reload admin with associated libraries before returning
	config.DB.Preload("Library").First(&admin, admin.ID)

	c.JSON(http.StatusCreated, gin.H{"message": "Admin registered successfully", "admin": admin})
}
