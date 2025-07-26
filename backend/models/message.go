package models

import "github.com/google/uuid"

type Message struct {
	Base
	ProjectID uuid.UUID
	Project   Project
	UserID    uuid.UUID
	User      User
	Role      string
	Content   string
}
