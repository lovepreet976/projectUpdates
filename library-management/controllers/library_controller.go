package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLibrary(c *gin.Context) {
	var input models.Library

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create library"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Library created successfully", "library": input})
}
