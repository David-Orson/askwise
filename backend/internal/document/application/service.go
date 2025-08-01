package application

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"
	"askwise.com/m/v2/internal/document/ports"
	"askwise.com/m/v2/internal/events"
	"github.com/google/uuid"
)

type DocumentService struct {
	Repo ports.DocumentRepository
	Bus  ports.EventBus
}

func NewDocumentService(repo ports.DocumentRepository, bus ports.EventBus) *DocumentService {
	return &DocumentService{Repo: repo, Bus: bus}
}

func (s *DocumentService) UploadDocument(ctx context.Context, projectID, userID uuid.UUID, fileName string) (*domain.Document, error) {
	doc, err := domain.NewDocument(projectID, userID, fileName)
	if err != nil {
		return nil, err
	}

	if err := s.Repo.Save(ctx, doc); err != nil {
		return nil, err
	}

	evt := events.DocumentUploadedEvent{
		DocumentID: doc.ID(),
		ProjectID:  projectID,
		UserID:     userID,
		FileName:   fileName,
		UploadedAt: doc.FormattedCreatedAt(),
	}

	if err := s.Bus.Publish(ctx, events.EventDocumentUploaded, evt); err != nil {
		return nil, err
	}

	return doc, nil
}
