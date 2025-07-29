package adapters

import (
	"askwise.com/m/v2/internal/shared/models"
)

type UserRecord struct {
	models.GormBase

	GoogleID string
	Name     string
	Email    string
	ImageURL string
}
