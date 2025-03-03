package main // Ensure it is "main" here

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password123" // Change this to your desired password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword)) // Print hashed password
}
