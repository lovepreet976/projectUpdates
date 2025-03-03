package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListLibraries(c *gin.Context) {
	var libraries []models.Library

	if err := config.DB.Find(&libraries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch libraries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"libraries": libraries})
}
