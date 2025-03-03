package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 🔍 Search Books by Title, Author, Publisher
func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	publisher := c.Query("publisher")

	var books []models.Book
	query := config.DB

	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("authors ILIKE ?", "%"+author+"%")
	}
	if publisher != "" {
		query = query.Where("publisher ILIKE ?", "%"+publisher+"%")
	}

	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}
