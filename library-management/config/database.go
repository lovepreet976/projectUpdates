package config

import (
	"fmt"
	"library-management/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=library_management sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-Migrate tables
	err = database.AutoMigrate(
		&models.Library{},
		&models.User{},
		&models.Book{},
		&models.RequestEvent{},
		&models.IssueRegistry{},
		&models.UserLibrary{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
	fmt.Println("Database connected and migrated successfully!")
}
