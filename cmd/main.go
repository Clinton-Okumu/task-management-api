package main

import (
	"log"
	"task-management-api/internal/auth"
	"task-management-api/internal/task"
	"task-management-api/internal/user"
	"task-management-api/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the database connection
	db.ConnectDB()

	// Optional: Perform migrations (create or update database tables)
	db.MigrateDB()

	// Initialize the Fiber app
	app := fiber.New()

	// Register routes
	app.Post("/auth/login", auth.Login)    // POST login
	app.Get("/tasks", task.GetAllTasks)    // GET all tasks
	app.Post("/tasks", task.CreateTask)    // POST create task
	app.Get("/users", user.GetUserProfile) // GET user profile

	// Start the server
	log.Println("Server is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
