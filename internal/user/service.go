package user

import (
	"errors"
	"task-management-api/internal/models"

	"gorm.io/gorm"
)

// GetUser retrieves a user by ID
func GetUser(db *gorm.DB, userID uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user's information
func UpdateUser(db *gorm.DB, userID uint, updates map[string]interface{}) error {
	var user models.User

	// Find the user
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Update the user
	if err := db.Model(&user).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}
