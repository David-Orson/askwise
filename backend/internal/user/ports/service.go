package ports

import (
	"context"

	"askwise.com/m/v2/internal/user/domain"
)

type UserService interface {
	Sync(ctx context.Context, googleID, name, email, image string) (*domain.User, error)
}
