package handlers

import (
	"errors"
	"strings"

	"askwise.com/m/v2/models"
	"askwise.com/m/v2/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) SyncUser(c *fiber.Ctx) error {
	// INFO: if we add another provider we will need to update this and the request such that it can identify the provider

	googleID := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Image string `json:"image"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	var user models.User
	result := h.DB.Where("google_id = ?", googleID).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user = models.User{
			Base:     models.NewBase(),
			GoogleID: googleID,
			Name:     body.Name,
			Email:    body.Email,
			ImageURL: body.Image,
		}

		if err := h.DB.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				return c.Status(fiber.StatusConflict).JSON(utils.Error("User already exists with this email or Google ID"))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(utils.Error("Failed to create user"))
		}
	} else if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Error("Database error"))
	}

	return c.JSON(user)
}
