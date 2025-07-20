package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || pass == "" || name == "" || port == "" {
		log.Fatalf("Missing one or more DB env vars: host=%s user=%s dbname=%s port=%s", host, user, name, port)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)

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

	if os.Getenv("ENV") == "development" {
		_ = godotenv.Load()
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// db := connectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("AskWise backend running.")
	})

	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong from backend 🚀")
	})

	/*
		app.Get("/api/db-ping", func(c *fiber.Ctx) error {
			sqlDB, _ := db.DB()
			err := sqlDB.Ping()
			if err != nil {
				return c.Status(500).SendString("DB connection failed: " + err.Error())
			}
			return c.SendString("Connected to Postgres ✅")
		})
	*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
