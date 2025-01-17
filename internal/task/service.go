package task

import (
	"errors"

	"gorm.io/gorm"
)

// CreateTask creates a new task in the database
func CreateTask(db *gorm.DB, task *Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.UserID == 0 {
		return errors.New("user ID is required")
	}

	// Save task to the database
	if err := db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// GetTasks retrieves all tasks for a specific user
func GetTasks(db *gorm.DB, userID uint) ([]Task, error) {
	var tasks []Task
	if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// UpdateTask updates an existing task
func UpdateTask(db *gorm.DB, taskID uint, updates map[string]interface{}) error {
	var task Task

	// Find the task by ID
	if err := db.First(&task, taskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}

	// Apply updates
	if err := db.Model(&task).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTask deletes a task by ID
func DeleteTask(db *gorm.DB, taskID uint) error {
	if err := db.Delete(&Task{}, taskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}
	return nil
}
