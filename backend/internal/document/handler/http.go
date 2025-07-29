package handler

import (
	"fmt"

	"askwise.com/m/v2/internal/document/ports"
	"askwise.com/m/v2/internal/document/presentation"
	"askwise.com/m/v2/internal/shared/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DocumentHandler struct {
	Service ports.DocumentService
}

func NewDocumentHandler(svc ports.DocumentService) *DocumentHandler {
	return &DocumentHandler{Service: svc}
}

func (h *DocumentHandler) Upload(c *fiber.Ctx) error {
	userIDString, ok := c.Locals("userID").(string)
	if !ok || userIDString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("unauthenticated"))
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Error("Invalid userID"))
	}

	projectID, err := uuid.Parse(c.Params("projectID"))
	if err != nil {
		return c.Status(400).JSON(utils.Error("Invalid projectID"))
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Error("Missing file"))
	}

	fileName := fileHeader.Filename

	doc, err := h.Service.UploadDocument(c.Context(), projectID, userID, fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Error(err.Error()))
	}

	fmt.Println("Document uploaded successfully:", doc.ID())

	return c.Status(fiber.StatusOK).JSON(utils.Success(presentation.FromDomain(doc)))
}
