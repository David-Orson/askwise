package adapters

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"
	"askwise.com/m/v2/internal/document/ports"
	"gorm.io/gorm"
)

type PostgresDocumentRepository struct {
	db *gorm.DB
}

func NewPostgresDocumentRepository(db *gorm.DB) ports.DocumentRepository {
	return &PostgresDocumentRepository{db: db}
}

func (r *PostgresDocumentRepository) Save(ctx context.Context, doc *domain.Document) error {
	return r.db.WithContext(ctx).Create(doc).Error
}

func (r *PostgresDocumentRepository) FindByID(ctx context.Context, id string) (*domain.Document, error) {
	var doc domain.Document
	if err := r.db.WithContext(ctx).First(&doc, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}
