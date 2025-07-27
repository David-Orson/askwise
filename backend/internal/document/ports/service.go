package ports

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"
	"github.com/google/uuid"
)

type DocumentService interface {
	UploadDocument(ctx context.Context, projectID, userID uuid.UUID, fileName string) (*domain.Document, error)
}
