package application

import (
	"context"

	"askwise.com/m/v2/internal/document/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(ctx context.Context, doc *domain.Document) error {
	args := m.Called(ctx, doc)
	return args.Error(0)
}

func (m *MockRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Document), args.Error(1)
}

type MockBus struct {
	mock.Mock
}

func (m *MockBus) Publish(ctx context.Context, topic string, payload any) error {
	args := m.Called(ctx, topic, payload)
	return args.Error(0)
}
