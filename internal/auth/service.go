package auth

import (
	"errors"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

// DB is the database instance for the auth package
var DB *gorm.DB

// Init initializes the database instance for the auth package
func Init(db *gorm.DB) {
	DB = db
}

// RegisterUser registers a new user
func RegisterUser(username, password, email string) error {
	// Check if username or email exists
	var existingUser User
	result := DB.Where("username = ? OR email = ?", username, email).First(&existingUser)
	if result.RowsAffected > 0 {
		return errors.New("username or email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user
	user := User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Save the user in the database
	if err := DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// LoginUser verifies user credentials
func LoginUser(username, password string) error {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		return errors.New("username is incorrect")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
