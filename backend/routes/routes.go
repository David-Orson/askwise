package routes

import (
	"askwise.com/m/v2/handlers"
	"askwise.com/m/v2/middleware"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *gorm.DB) {
	auth := app.Group("/auth")
	api := app.Group("/api", middleware.AuthRequired())
	public := app.Group("/public")

	// Auth routes
	RegisterAuthRoutes(auth, db)

	// Authenticated routes
	RegisterProjectRoutes(api, db)

	// Public routes
	RegisterPublicRoutes(public, db)
}

func RegisterAuthRoutes(router fiber.Router, db *gorm.DB) {
	h := &handlers.AuthHandler{DB: db}

	router.Post("/sync", h.SyncUser)
}

func RegisterProjectRoutes(router fiber.Router, db *gorm.DB) {
	h := &handlers.ProjectHandler{DB: db}

	router.Post("/projects", h.CreateProject)
	router.Get("/projects", h.ListProjects)
}

func RegisterPublicRoutes(router fiber.Router, db *gorm.DB) {
	h := &handlers.PublicHandler{DB: db}

	router.Get("/ping", h.Ping)
	router.Get("/", h.HealthCheck)
	router.Get("/healthcheck", h.HealthCheck)
	router.Get("/db-ping", h.DBPing)
}
