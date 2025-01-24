package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database instance
var dbPool *pgxpool.Pool

// ConnectDB initializes the PostgreSQL connection using pgx
func ConnectDB() {
	// Database connection string (adjust according to your setup)
	// Example: "postgres://username:password@localhost:5432/dbname"
	connString := "postgres://udummy:kujwhehfuw32433@localhost:5432/db_udummy"

	// Create a new context with timeout
	ctx := context.Background()

	// Initialize connection pool
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Test the connection
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}

	log.Println("Connected to the database successfully!")

	dbPool = pool
}

// CloseDB closes the database connection pool
func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
		log.Println("Database connection closed.")
	}
}

func main() {
	// Connect to the database
	ConnectDB()
	defer CloseDB()

	// Initialize Fiber app
	app := fiber.New()

	// Simple route to test database connection
	app.Get("/", func(c *fiber.Ctx) error {
		ctx := context.Background()
		var currentTime string

		// Execute a simple query
		err := dbPool.QueryRow(ctx, "SELECT NOW()").Scan(&currentTime)
		if err != nil {
			log.Printf("Query failed: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to query database.")
		}

		return c.SendString("Current time from database: " + currentTime)
	})

	// Start the Fiber app
	log.Fatal(app.Listen(":3000"))
}
