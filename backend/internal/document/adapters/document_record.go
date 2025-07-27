package adapters

import (
	"askwise.com/m/v2/internal/shared/models"
	"github.com/google/uuid"
)

type DocumentRecord struct {
	models.GormBase

	ProjectID uuid.UUID
	UserID    uuid.UUID
	FileName  string
}
