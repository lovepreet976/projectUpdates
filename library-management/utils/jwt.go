package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key for signing JWT tokens
var jwtKey = []byte("secret_key")

// GenerateJWT creates a JWT token for a user
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expires in 1 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // FIX: Correct JWT creation
	return token.SignedString(jwtKey)
}

// ValidateJWT parses and validates a JWT token
func ValidateJWT(tokenString string) (uint, string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	// Extract userID and role from claims correctly
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", errors.New("invalid user_id")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", errors.New("invalid role")
	}

	return uint(userIDFloat), role, nil
}
