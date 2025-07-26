package models

import "github.com/google/uuid"

type Document struct {
	Base
	Name      string
	ProjectID uuid.UUID
	Project   Project
	UserID    uuid.UUID
	User      User
	FileURL   string
}
