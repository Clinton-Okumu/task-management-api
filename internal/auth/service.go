package auth

import (
	"errors"
	"task-management-api/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user
func RegisterUser(username, password, email string) error {
	// check if username or email exists
	var existingUser User
	result := db.DB.Where("username = ? OR email = ?", username, email).First(&existingUser)
	if results.RowsAffected > 0 {
		return errors.New("username or email already exists")
	}

	// hashing password
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

	// save user in the database
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func LoginUser(username, password string) error {
	var user User
	result := db.DB.Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 {
		return errors.New("username or username is incorrect")
	}

	// compare hashed password
	if err != bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
