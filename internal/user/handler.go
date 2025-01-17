package user

import (
	"task-management-api/pkg/db"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers user-related routes
func RegisterRoutes(app *fiber.App) {
	app.Get("/users/:id", GetUserHandler)
	app.Put("/users/:id", UpdateUserHandler)
}

// GetUserHandler handles retrieving a user profile
func GetUserHandler(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := GetUser(db.DB, uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

// UpdateUserHandler handles updating a user profile
func UpdateUserHandler(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := UpdateUser(db.DB, uint(userID), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
