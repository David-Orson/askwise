package ports

import (
	"context"

	"askwise.com/m/v2/internal/user/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByGoogleID(ctx context.Context, googleID string) (*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
}
