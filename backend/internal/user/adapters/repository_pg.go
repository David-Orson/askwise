package adapters

import (
	"context"

	"askwise.com/m/v2/internal/user/domain"
	"askwise.com/m/v2/internal/user/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) ports.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var record UserRecord
	if err := r.db.WithContext(ctx).First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}

	user := toDomain(&record)

	return user, nil
}

func (r *PostgresUserRepository) FindByGoogleID(ctx context.Context, googleID string) (*domain.User, error) {
	var record UserRecord
	if err := r.db.WithContext(ctx).First(&record, "google_id = ?", googleID).Error; err != nil {
		return nil, err
	}

	user := toDomain(&record)

	return user, nil
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *domain.User) error {
	record := toRecord(user)
	return r.db.WithContext(ctx).Create(record).Error
}
