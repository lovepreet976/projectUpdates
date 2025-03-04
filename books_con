package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 🟢 Add Book (Create or Increment Copies)
func AddBook(c *gin.Context) {
	var input models.Book

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure book has valid copies
	if input.TotalCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number of copies must be greater than zero"})
		return
	}

	var existingBook models.Book
	if err := config.DB.Where("isbn = ? AND library_id = ?", input.ISBN, input.LibraryID).First(&existingBook).Error; err == nil {
		// 📌 Book exists → Increment copies
		existingBook.TotalCopies += input.TotalCopies
		existingBook.AvailableCopies += input.TotalCopies

		if err := config.DB.Save(&existingBook).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book copies"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book copies updated successfully", "book": existingBook})
		return
	}

	// 📌 New Book → Insert into DB
	input.AvailableCopies = input.TotalCopies
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully", "book": input})
}

// 🟡 Update Book Details
func UpdateBook(c *gin.Context) {
	isbn := c.Param("isbn")
	var book models.Book

	// Check if book exists
	if err := config.DB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prevent reducing total copies below issued copies
	issuedCopies := book.TotalCopies - book.AvailableCopies
	if input.TotalCopies < issuedCopies {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total copies cannot be less than issued copies"})
		return
	}

	// ✅ Update Book Details
	book.Title = input.Title
	book.Authors = input.Authors
	book.Publisher = input.Publisher
	book.Version = input.Version
	book.TotalCopies = input.TotalCopies
	book.AvailableCopies = input.TotalCopies - issuedCopies // Ensure AvailableCopies is adjusted properly

	if err := config.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "book": book})
}

// 🔴 Remove a Book (Only if copies exist)
func RemoveBook(c *gin.Context) {
	isbn := c.Param("isbn")
	var book models.Book

	// Find the book
	if err := config.DB.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// 📌 Prevent deletion if any copies are issued
	issuedCopies := book.TotalCopies - book.AvailableCopies
	if issuedCopies > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot remove book as some copies are issued"})
		return
	}

	// 📌 Decrement available copies
	if book.TotalCopies > 1 {
		book.TotalCopies--
		book.AvailableCopies--
		if err := config.DB.Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrement book copies"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book copies decremented", "book": book})
	} else {
		// 📌 Delete book if no copies remain
		if err := config.DB.Delete(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book removed from inventory"})
	}
}
