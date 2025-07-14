package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	maxRetries := 10
	for retries := 1; retries <= maxRetries; retries++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to Postgres!")
			return db
		}

		log.Printf("Failed to connect to DB (attempt %d/%d): %v\n", retries, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Could not connect to DB after retries:", err)
	return nil
}

func main() {
	app := fiber.New()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found — falling back to environment variables")
	}

	db := connectDB()

	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong from backend 🚀")
	})

	app.Get("/api/db-ping", func(c *fiber.Ctx) error {
		sqlDB, _ := db.DB()
		err := sqlDB.Ping()
		if err != nil {
			return c.Status(500).SendString("DB connection failed: " + err.Error())
		}
		return c.SendString("Connected to Postgres ✅")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
