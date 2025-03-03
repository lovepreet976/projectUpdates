package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 📄 List Issue Requests
func ListIssueRequests(c *gin.Context) {
	var requests []models.RequestEvent
	if err := config.DB.Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch issue requests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

// ✅ Approve Issue Request
func ApproveIssue(c *gin.Context) {
	requestID := c.Param("id")
	var request models.RequestEvent

	if err := config.DB.First(&request, requestID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue request not found"})
		return
	}

	if request.ApprovalDate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request is already approved"})
		return
	}

	var book models.Book
	if err := config.DB.Where("isbn = ?", request.BookID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.AvailableCopies == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No available copies to issue"})
		return
	}

	// 📌 Decrease available copies and mark as issued
	book.AvailableCopies--
	config.DB.Save(&book)

	// ✅ Mark the request as approved
	now := time.Now().Unix()
	request.ApprovalDate = &now

	if err := config.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not approve request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue request approved", "request": request})
}

// ❌ Disapprove Issue Request
func DisapproveIssue(c *gin.Context) {
	requestID := c.Param("id")
	var request models.RequestEvent

	if err := config.DB.First(&request, requestID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue request not found"})
		return
	}

	config.DB.Delete(&request)
	c.JSON(http.StatusOK, gin.H{"message": "Issue request disapproved"})
}

// 🔄 Issue Book to a User
func IssueBookToUser(c *gin.Context) {
	isbn := c.Param("isbn")

	// ✅ Extract Admin ID from JWT
	adminID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
		return
	}

	// ✅ Parse JSON Input
	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// ✅ Ensure Book Exists
	var book models.Book
	if err := config.DB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// ✅ Check if Book is Available
	if book.AvailableCopies == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No available copies to issue"})
		return
	}

	// ✅ Reduce Available Copies
	book.AvailableCopies--
	config.DB.Save(&book)

	// ✅ Record the Issue in `IssueRegistry`
	issueRecord := models.IssueRegistry{
		ISBN:               isbn,
		ReaderID:           input.UserID,
		IssueApproverID:    adminID.(uint),
		IssueStatus:        "issued",
		IssueDate:          time.Now().Unix(),
		ExpectedReturnDate: time.Now().AddDate(0, 0, 14).Unix(), // Default return date in 2 weeks
		ReturnDate:         0,
		ReturnApproverID:   0,
	}

	if err := config.DB.Create(&issueRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not issue book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book issued successfully", "issue": issueRecord})
}
