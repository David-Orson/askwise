package domain

import (
	"errors"
	"time"

	"askwise.com/m/v2/internal/shared/models"
	"github.com/google/uuid"
)

type Document struct {
	Base      models.DomainBase
	projectID uuid.UUID
	userID    uuid.UUID
	fileName  string
}

func NewDocument(projectID, userID uuid.UUID, fileName string) (*Document, error) {
	if projectID == uuid.Nil {
		return nil, errors.New("projectID must be a valid UUID")
	}
	if userID == uuid.Nil {
		return nil, errors.New("userID must be a valid UUID")
	}
	if err := ValidateFileName(fileName); err != nil {
		return nil, err
	}

	return &Document{
		Base:      models.NewDomainBase(),
		projectID: projectID,
		userID:    userID,
		fileName:  fileName,
	}, nil
}

func (d *Document) ID() uuid.UUID {
	return d.Base.ID
}

func (d *Document) ProjectID() uuid.UUID {
	return d.projectID
}

func (d *Document) UserID() uuid.UUID {
	return d.userID
}

func (d *Document) FileName() string {
	return d.fileName
}

func (d *Document) CreatedAt() time.Time {
	return d.Base.CreatedAt
}

func (d *Document) UpdatedAt() time.Time {
	return d.Base.UpdatedAt
}

func (d *Document) FormattedCreatedAt() string {
	return d.CreatedAt().Format(time.RFC3339)
}

func (d *Document) FormattedUpdatedAt() string {
	return d.UpdatedAt().Format(time.RFC3339)
}

func ReconstructDocument(
	id, projectID, userID uuid.UUID,
	fileName string,
	created, updated time.Time,
) *Document {
	return &Document{
		Base: models.DomainBase{
			ID:        id,
			CreatedAt: created,
			UpdatedAt: updated,
		},
		projectID: projectID,
		userID:    userID,
		fileName:  fileName,
	}
}

func NewTestDocument(id uuid.UUID, projectID, userID uuid.UUID, fileName string, created time.Time) *Document {
	return &Document{
		Base: models.DomainBase{
			ID:        id,
			CreatedAt: created,
			UpdatedAt: created,
		},
		projectID: projectID,
		userID:    userID,
		fileName:  fileName,
	}
}
