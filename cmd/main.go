package main

import (
	"log"
	"task-management-api/internal/auth"
	"task-management-api/internal/task"
	"task-management-api/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the database connection
	db.ConnectDB()

	// Inject the database into packages
	auth.Init(db.DB)
	task.Init(db.DB)

	// Perform migrations
	if err := db.MigrateDB(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize the Fiber app
	app := fiber.New()

	// Register routes
	auth.RegisterRoutes(app)
	task.RegisterRoutes(app)

	// Start the server
	log.Println("Server is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
