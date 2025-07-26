package domain

import (
	"time"

	"askwise.com/m/v2/internal/shared/models"
	"github.com/google/uuid"
)

type Document struct {
	models.Base
	projectID string
	userID    string
	fileName  string
}

func NewDocument(projectID, userID, fileName string) *Document {
	return &Document{
		Base:      models.NewBase(),
		projectID: projectID,
		userID:    userID,
		fileName:  fileName,
	}
}

func (d *Document) ID() uuid.UUID {
	return d.Base.ID
}

func (d *Document) ProjectID() string {
	return d.projectID
}

func (d *Document) UserID() string {
	return d.userID
}

func (d *Document) FileName() string {
	return d.fileName
}

func (d *Document) createdAt() time.Time {
	return d.Base.CreatedAt
}

func (d *Document) FormattedCreatedAt() string {
	return d.createdAt().Format(time.RFC3339)
}
