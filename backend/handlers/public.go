package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PublicHandler struct {
	DB *gorm.DB
}

func (h *PublicHandler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong from backend 🚀")
}

func (h *PublicHandler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("AskWise backend running.")
}

func (h *PublicHandler) DBPing(c *fiber.Ctx) error {
	sqlDB, _ := h.DB.DB()
	err := sqlDB.Ping()
	if err != nil {
		return c.Status(500).SendString("DB connection failed: " + err.Error())
	}
	return c.SendString("Connected to Postgres ✅")
}
