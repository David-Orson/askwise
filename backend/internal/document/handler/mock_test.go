package handler

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockDocService struct {
	mock.Mock
}

func (m *MockDocService) UploadDocument(ctx context.Context, projectID, userID uuid.UUID, fileName string) (*domain.Document, error) {
	args := m.Called(ctx, projectID, userID, fileName)
	if doc, ok := args.Get(0).(*domain.Document); ok {
		return doc, args.Error(1)
	}
	return nil, args.Error(1)
}
