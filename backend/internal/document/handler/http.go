package handler

import (
	"askwise.com/m/v2/internal/document/application"
	"askwise.com/m/v2/internal/document/presentation"
	"askwise.com/m/v2/internal/shared/utils"
	"github.com/gofiber/fiber/v2"
)

type DocumentHandler struct {
	Service *application.DocumentService
}

func NewDocumentHandler(svc *application.DocumentService) *DocumentHandler {
	return &DocumentHandler{Service: svc}
}

func (h *DocumentHandler) Upload(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("unauthenticated"))
	}
	projectID := c.Params("projectID")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Error("Missing file"))
	}

	fileName := fileHeader.Filename

	doc, err := h.Service.UploadDocument(c.Context(), projectID, userID, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(utils.Success(presentation.FromDomain(doc)))
}
