package models

import "github.com/google/uuid"

type Project struct {
	Base
	Name      string
	UserID    uuid.UUID
	User      User
	Documents []Document
}
