package handlers

import (
	"askwise.com/m/v2/models"
	"askwise.com/m/v2/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DebugUserID = "95f85564-d4b8-4baa-8351-741073e13203"

type ProjectHandler struct {
	DB *gorm.DB
}

// POST /api/projects
func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	var body request

	if err := c.BodyParser(&body); err != nil || body.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Error("Project name is required"))
	}

	uid, err := uuid.Parse(DebugUserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Error("Invalid user ID"))
	}

	project := models.Project{
		Base:   models.NewBase(),
		Name:   body.Name,
		UserID: uid,
	}

	if err := h.DB.Create(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Error("Failed to create project"))
	}

	return c.JSON(utils.Success(project))
}

// GET /api/projects
func (h *ProjectHandler) ListProjects(c *fiber.Ctx) error {
	var projects []models.Project

	if err := h.DB.Where("user_id = ?", DebugUserID).Find(&projects).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Error("Failed to fetch projects"))
	}

	return c.JSON(utils.Success(projects))
}
