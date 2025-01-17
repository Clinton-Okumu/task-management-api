package auth

import (
	"errors"
	"task-management-api/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterUser registers a new user
func RegisterUser(db *gorm.DB, username, password, email string) error {
	// Check if username or email exists
	var existingUser models.User
	result := db.Where("username = ? OR email = ?", username, email).First(&existingUser)
	if result.RowsAffected > 0 {
		return errors.New("username or email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user
	user := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Save the user in the database
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// LoginUser verifies user credentials
func LoginUser(db *gorm.DB, username, password string) error {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		return errors.New("username is incorrect")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
