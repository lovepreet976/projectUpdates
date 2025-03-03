package main

import (
	"library-management/config"
	"library-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()       // Create a new Gin router
	config.ConnectDatabase() // Initialize the database

	r = routes.SetupRouter() // Assign the configured router
	r.Run(":8080")           // Start the server
}
