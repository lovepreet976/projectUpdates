package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestIssue handles book issue requests by users
func RequestIssue(c *gin.Context) {
	var input struct {
		BookID string `json:"isbn" binding:"required"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user ID from JWT
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
		return
	}

	// Validate book existence
	var book models.Book
	if err := config.DB.Where("isbn = ?", input.BookID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Ensure book is available
	if book.AvailableCopies == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available for issue"})
		return
	}

	// ✅ Check if the user already has a pending request for this book
	var existingRequest models.RequestEvent
	if err := config.DB.Where("reader_id = ? AND book_id = ? AND approval_date IS NULL", userID, input.BookID).First(&existingRequest).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You already have a pending request for this book"})
		return
	}

	// ✅ Create new issue request
	request := models.RequestEvent{
		BookID:       input.BookID,
		ReaderID:     userID.(uint),
		RequestDate:  time.Now().Unix(),
		ApprovalDate: nil, // Not approved yet
		ApproverID:   nil, // No admin assigned yet
		RequestType:  "issue",
	}

	// ✅ Save request to DB
	if err := config.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create issue request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Issue request submitted", "request": request})
}
