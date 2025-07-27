package models

import (
	"time"

	"github.com/google/uuid"
)

type DomainBase struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDomainBase() DomainBase {
	return DomainBase{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
