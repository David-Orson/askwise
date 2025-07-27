package adapters

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"
	"askwise.com/m/v2/internal/document/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresDocumentRepository struct {
	db *gorm.DB
}

func NewPostgresDocumentRepository(db *gorm.DB) ports.DocumentRepository {
	return &PostgresDocumentRepository{db: db}
}

func (r *PostgresDocumentRepository) Save(ctx context.Context, doc *domain.Document) error {
	record := toRecord(doc)
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *PostgresDocumentRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	var record DocumentRecord
	if err := r.db.WithContext(ctx).First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}

	doc := toDomain(&record)

	return doc, nil
}
