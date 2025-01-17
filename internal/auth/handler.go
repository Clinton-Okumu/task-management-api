package auth

import (
	"task-management-api/pkg/db"

	"github.com/gofiber/fiber/v2"
)

// RegisterHandler handles user registration
func RegisterHandler(c *fiber.Ctx) error {
	// Parse the request body
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Call the service layer
	if err := RegisterUser(db.DB, req.Username, req.Password, req.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// LoginHandler handles user login
func LoginHandler(c *fiber.Ctx) error {
	// Parse the request body
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Call the service layer
	if err := LoginUser(db.DB, req.Username, req.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}
