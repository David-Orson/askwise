package adapters

import (
	"askwise.com/m/v2/internal/shared/models"
	"askwise.com/m/v2/internal/user/domain"
)

func toRecord(u *domain.User) *UserRecord {
	return &UserRecord{
		GormBase: models.GormBase{
			ID:        u.ID(),
			CreatedAt: u.CreatedAt(),
			UpdatedAt: u.UpdatedAt(),
		},
		GoogleID: u.GoogleID(),
		Name:     u.Name(),
		Email:    u.Email(),
		ImageURL: u.ImageURL(),
	}
}

func toDomain(r *UserRecord) *domain.User {
	return domain.ReconstructUser(
		r.ID,
		r.GoogleID,
		r.Name,
		r.Email,
		r.ImageURL,
		r.CreatedAt,
		r.UpdatedAt,
	)
}
