package handler

import (
	"strings"

	"askwise.com/m/v2/internal/shared/utils"
	"askwise.com/m/v2/internal/user/ports"
	"askwise.com/m/v2/internal/user/presentation"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service ports.UserService
}

func NewUserHandler(svc ports.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

func (h *UserHandler) Sync(c *fiber.Ctx) error {
	googleID := extractBearerToken(c)
	if googleID == "" {
		return c.Status(401).JSON(utils.Error("Missing Google ID"))
	}

	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Image string `json:"image"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(utils.Error("Invalid request body"))
	}

	user, err := h.Service.Sync(c.Context(), googleID, body.Name, body.Email, body.Image)
	if err != nil {
		return c.Status(500).JSON(utils.Error("Sync failed: " + err.Error()))
	}

	return c.JSON(presentation.FromDomain(user))
}

func extractBearerToken(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
