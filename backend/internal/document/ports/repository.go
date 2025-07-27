package ports

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"
	"github.com/google/uuid"
)

type DocumentRepository interface {
	Save(ctx context.Context, doc *domain.Document) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Document, error)
}
