package cmd

import (
	"fmt"
	"log"
	"os"

	"askwise.com/m/v2/internal/document/adapters"
	"askwise.com/m/v2/internal/document/application"
	"askwise.com/m/v2/internal/document/handler"
	"askwise.com/m/v2/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	redisAdapter "askwise.com/m/v2/internal/events/redis"
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

	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // TODO: add cloud endpoint
		Password: "",               // TODO: add password
		DB:       0,
	})

	docRepo := adapters.NewPostgresDocumentRepository(db)
	eventBus := redisAdapter.NewRedisEventBus(redis)

	docSvc := application.NewDocumentService(docRepo, eventBus)
	docHandler := handler.NewDocumentHandler(docSvc)

	app.Post("/api/projects/:projectID/upload", docHandler.Upload)

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
