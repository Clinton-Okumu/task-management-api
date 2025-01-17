package models

import "gorm.io/gorm"

// Task represents a task in the system
type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`     // Title of the task, required
	Description string `json:"description"`               // Description of the task
	Completed   bool   `gorm:"not null" json:"completed"` // Task completion status
	UserID      uint   `gorm:"not null" json:"user_id"`   // Foreign key for the associated user
}

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"` // Unique username, required
	Password string `gorm:"not null" json:"password"`        // User's hashed password
	Email    string `gorm:"unique;not null" json:"email"`    // Unique email, required
}
