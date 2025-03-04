# projectUpdates
🔹 Step 1: Owner Login
📌 Endpoint:
http
Copy code
POST /auth/login
📌 Request Body (JSON):
json
Copy code
{
    "email": "owner@example.com",
    "password": "securepassword"
}
📌 Expected Response:
json
Copy code
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
✅ Copy the JWT Token for the next requests.
________________________________________
🔹 Step 2: Create a Library (Owner Only)
📌 Endpoint:
http
Copy code
POST /api/library
📌 Headers:
plaintext
Copy code
Authorization: Bearer <OWNER_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "name": "Central Library"
}
📌 Expected Response:
json
Copy code
{
    "message": "Library created successfully",
    "library": {
        "id": 1,
        "name": "Central Library"
    }
}
✅ Copy the library_id for the next request.
________________________________________
🔹 Step 3: Create an Admin for the Library
📌 Endpoint:
http
Copy code
POST /api/admin
📌 Headers:
plaintext
Copy code
Authorization: Bearer <OWNER_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "email": "admin@example.com",
    "password": "securepassword",
    "role": "admin",
    "library_id": 1
}
📌 Expected Response:
json
Copy code
{
    "message": "Admin registered successfully",
    "admin": {
        "id": 2,
        "email": "admin@example.com",
        "role": "admin",
        "library_id": 1
    }
}
✅ Copy the admin_id for later.
________________________________________
🔹 Step 4: Admin Login
📌 Endpoint:
http
Copy code
POST /auth/login
📌 Request Body (JSON):
json
Copy code
{
    "email": "admin@example.com",
    "password": "securepassword"
}
📌 Expected Response:
json
Copy code
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
✅ Copy the JWT Token for Admin.
________________________________________
🔹 Step 5: Add a Book to the Library (Admin Only)
📌 Endpoint:
http
Copy code
POST /api/book
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "isbn": "978-3-16-148410-0",
    "title": "Golang Mastery",
    "author": "John Doe",
    "publisher": "Tech Books",
    "copies": 5,
    "library_id": 1
}
📌 Expected Response:
json
Copy code
{
    "message": "Book added successfully",
    "book": {
        "isbn": "978-3-16-148410-0",
        "title": "Golang Mastery",
        "author": "John Doe",
        "publisher": "Tech Books",
        "copies": 5,
        "library_id": 1
    }
}
________________________________________
🔹 Step 6: User Registration
📌 Endpoint:
http
Copy code
POST /api/user
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "email": "user@example.com",
    "password": "securepassword",
    "role": "user",
    "library_ids": [1]
}
📌 Expected Response:
json
Copy code
{
    "message": "User registered successfully",
    "user": {
        "id": 3,
        "email": "user@example.com",
        "role": "user",
        "library_ids": [1]
    }
}
✅ Copy the user_id.
________________________________________
🔹 Step 7: User Login
📌 Endpoint:
http
Copy code
POST /auth/login
📌 Request Body (JSON):
json
Copy code
{
    "email": "user@example.com",
    "password": "securepassword"
}
📌 Expected Response:
json
Copy code
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
✅ Copy the JWT Token for User.
________________________________________
🔹 Step 8: Search for Books (User Only)
📌 Endpoint:
http
Copy code
GET /api/books/search?title=Golang
📌 Headers:
plaintext
Copy code
Authorization: Bearer <USER_TOKEN>
📌 Expected Response:
json
Copy code
{
    "books": [
        {
            "isbn": "978-3-16-148410-0",
            "title": "Golang Mastery",
            "author": "John Doe",
            "publisher": "Tech Books",
            "copies": 5,
            "library_id": 1
        }
    ]
}
________________________________________
🔹 Step 9: Request to Issue a Book (User Only)
📌 Endpoint:
http
Copy code
POST /api/issue
📌 Headers:
plaintext
Copy code
Authorization: Bearer <USER_TOKEN>
Content-Type: application/json
📌 Request Body (JSON):
json
Copy code
{
    "user_id": 3,
    "isbn": "978-3-16-148410-0",
    "library_id": 1
}
📌 Expected Response:
json
Copy code
{
    "message": "Issue request submitted",
    "request": {
        "user_id": 3,
        "isbn": "978-3-16-148410-0",
        "library_id": 1,
        "request_date": 1700000000
    }
}
________________________________________
🔹 Step 10: Admin Approves the Request
📌 Endpoint:
http
Copy code
PUT /api/issue/approve/:id
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
📌 Expected Response:
json
Copy code
{
    "message": "Issue request approved",
    "issue": {
        "id": 1,
        "user_id": 3,
        "isbn": "978-3-16-148410-0",
        "status": "approved"
    }
}
________________________________________
🔹 Step 11: Admin Issues the Book to User
📌 Endpoint:
http
Copy code
POST /api/issue/book/:isbn
📌 Headers:
plaintext
Copy code
Authorization: Bearer <ADMIN_TOKEN>
📌 Expected Response:
json
Copy code
{
    "message": "Book issued successfully",
    "issue": {
        "id": 1,
        "user_id": 3,
        "isbn": "978-3-16-148410-0",
        "status": "issued"
    }
}



new updations>>>
package controllers

import (
	"library-management/config"
	"library-management/models"
	"net/http"
	"time"

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
		query = query.Where("authors ILIKE ?", "%"+author+"%") // Fix column name
	}
	if publisher != "" {
		query = query.Where("publisher ILIKE ?", "%"+publisher+"%")
	}

	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching books"})
		return
	}

	// ✅ Check availability and fetch next available date if needed
	var response []gin.H
	for _, book := range books {
		bookData := gin.H{
			"isbn":             book.ISBN,
			"title":            book.Title,
			"author":           book.Authors, // Fix column name
			"publisher":        book.Publisher,
			"available_copies": book.AvailableCopies,
		}

		if book.AvailableCopies == 0 {
			var nextAvailableDate time.Time
			var issue models.IssueRegistry

			// Find the earliest expected return date
			if err := config.DB.Where("isbn = ?", book.ISBN).
				Order("expected_return_date ASC").
				First(&issue).Error; err == nil {
				nextAvailableDate = time.Unix(issue.ExpectedReturnDate, 0)
				bookData["next_available_date"] = nextAvailableDate.Format("2006-01-02 15:04:05")
			} else {
				bookData["next_available_date"] = "Unknown"
			}
		}

		response = append(response, bookData)
	}

	c.JSON(http.StatusOK, gin.H{"books": response})
}

issue,,

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

	// ✅ Just approve, don't reduce copies here
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

	// ✅ Set Issue & Return Dates
	issueDate := time.Now()
	expectedReturnDate := issueDate.AddDate(0, 0, 14) // 2 weeks later

	// ✅ Record the Issue in `IssueRegistry`
	issueRecord := models.IssueRegistry{
		ISBN:               isbn,
		ReaderID:           input.UserID,
		IssueApproverID:    adminID.(uint),
		IssueStatus:        "issued",
		IssueDate:          issueDate.Unix(),
		ExpectedReturnDate: expectedReturnDate.Unix(),
		ReturnDate:         0,
		ReturnApproverID:   0,
	}

	if err := config.DB.Create(&issueRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not issue book"})
		return
	}

	// ✅ Format Dates for Readable Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Book issued successfully",
		"issue": gin.H{
			"id":                   issueRecord.ID,
			"reader_id":            issueRecord.ReaderID,
			"isbn":                 issueRecord.ISBN,
			"issue_status":         issueRecord.IssueStatus,
			"issue_date":           issueDate.Format("2006-01-02 15:04:05"),
			"expected_return_date": expectedReturnDate.Format("2006-01-02 15:04:05"),
		},
	})
}



