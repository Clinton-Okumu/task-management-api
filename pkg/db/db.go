package db

import (
	"fmt"
	"log"
	"os"
	"task-management-api/internal/auth"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct DSN from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	// Open database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Get the underlying *sql.DB instance
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting DB instance: %v", err)
	}

	// Ping the database to confirm the connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to database")
}

// MigrateDB handles database migrations
func MigrateDB() error {
	err := DB.AutoMigrate(
		&auth.User{}, // User model
	)
	if err != nil {
		return fmt.Errorf("error migrating database: %w", err)
	}
	log.Println("Database migration completed")
	return nil
}
