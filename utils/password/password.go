package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash from a password string
func HashPassword(password string) (string, error) {
	// Cost of 12 is a good balance between security and performance
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword compares a password against a hash to check if they match
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
