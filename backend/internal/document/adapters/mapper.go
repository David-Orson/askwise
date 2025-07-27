package adapters

import (
	"askwise.com/m/v2/internal/document/domain"
	"askwise.com/m/v2/internal/shared/models"
)

func toRecord(d *domain.Document) *DocumentRecord {
	return &DocumentRecord{
		GormBase: models.GormBase{
			ID:        d.ID(),
			CreatedAt: d.CreatedAt(),
			UpdatedAt: d.UpdatedAt(),
		},
		ProjectID: d.ProjectID(),
		UserID:    d.UserID(),
		FileName:  d.FileName(),
	}
}

func toDomain(r *DocumentRecord) *domain.Document {
	return domain.ReconstructDocument(
		r.ID,
		r.ProjectID,
		r.UserID,
		r.FileName,
		r.CreatedAt,
		r.UpdatedAt,
	)
}
