package routes

import (
	"library-management/controllers"
	"library-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes (No authentication required)
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	// Protected API routes (Require authentication)
	api := r.Group("/api")

	{
		r.GET("/libraries", controllers.ListLibraries)

		// Owner-Only Routes
		ownerRoutes := api.Group("").Use(middleware.AuthMiddleware("owner"))
		{
			ownerRoutes.POST("/library", controllers.CreateLibrary)  // ✅ Owner can create a library
			ownerRoutes.POST("/admin", controllers.RegisterAdmin)    // ✅ Owner can create Admins
			ownerRoutes.POST("/owner", controllers.RegisterOwnerNew) // ✅ Owner can create a new Owner (Fixed Route)
		}

		// Admin-Only Routes
		adminRoutes := api.Group("").Use(middleware.AuthMiddleware("admin"))
		{
			adminRoutes.POST("/user", controllers.RegisterUser)

			// 📚 Book Management
			adminRoutes.POST("/book", controllers.AddBook)            // ✅ Admin can add books
			adminRoutes.PUT("/book/:isbn", controllers.UpdateBook)    // ✅ Admin can update book details (copies, title, etc.)
			adminRoutes.DELETE("/book/:isbn", controllers.RemoveBook) // ✅ Admin can remove books

			// 📄 Issue Request Management
			adminRoutes.GET("/issues", controllers.ListIssueRequests)             // ✅ Admin can list issue requests
			adminRoutes.PUT("/issue/approve/:id", controllers.ApproveIssue)       // ✅ Admin can approve issue requests
			adminRoutes.PUT("/issue/disapprove/:id", controllers.DisapproveIssue) // ✅ Admin can disapprove issue requests

			// 📖 Issue Books to Users
			adminRoutes.POST("/issue/book/:isbn", controllers.IssueBookToUser) // ✅ Admin can issue books to a reader

		}

		// User-Only Routes
		userRoutes := api.Group("").Use(middleware.AuthMiddleware("user"))
		{
			// 📌 Book Search
			userRoutes.GET("/books/search", controllers.SearchBooks) // ✅ Users can search books by title, author, publisher

			// 📄 Request a Book
			userRoutes.POST("/issue", controllers.RequestIssue) // ✅ Users can request book issues
		}
	}

	return r
}
