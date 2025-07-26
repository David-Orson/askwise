package main

import (
	"fmt"
	"log"
	"os"

	"askwise.com/m/v2/routes"
	"askwise.com/m/v2/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	_ = godotenv.Load()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://askwise.vercel.app",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	db := utils.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.Register(app, db)

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
