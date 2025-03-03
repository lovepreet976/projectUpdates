package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password, avoiding double hashing.
func HashPassword(password string) (string, error) {
	if strings.HasPrefix(password, "$2a$") { // bcrypt hashes start with "$2a$"
		return password, nil // Already hashed, return as is
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with a plain-text password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
