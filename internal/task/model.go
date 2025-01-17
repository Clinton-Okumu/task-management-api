package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Completed   bool   `gorm:"not null" json:"completed"`
	UserID      uint   `gorm:"not null" json:"user_id"`
}
