package task

import (
	"task-management-api/pkg/db"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers task-related routes
func RegisterRoutes(app *fiber.App) {
	app.Post("/tasks", CreateTaskHandler)
	app.Get("/tasks", GetTasksHandler)
	app.Put("/tasks/:id", UpdateTaskHandler)
	app.Delete("/tasks/:id", DeleteTaskHandler)
}

// CreateTaskHandler handles creating a new task
func CreateTaskHandler(c *fiber.Ctx) error {
	var task Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := CreateTask(db.DB, &task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

// GetTasksHandler handles retrieving all tasks for a user
func GetTasksHandler(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint) // Assuming middleware sets this
	tasks, err := GetTasks(db.DB, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}

// UpdateTaskHandler handles updating a task
func UpdateTaskHandler(c *fiber.Ctx) error {
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := UpdateTask(db.DB, uint(taskID), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteTaskHandler handles deleting a task
func DeleteTaskHandler(c *fiber.Ctx) error {
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	if err := DeleteTask(db.DB, uint(taskID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
