package presentation

import (
	"askwise.com/m/v2/internal/document/domain"
)

type DocumentResponse struct {
	ID        string `json:"id"`
	FileName  string `json:"file_name"`
	CreatedAt string `json:"created"`
}

func FromDomain(doc *domain.Document) DocumentResponse {
	return DocumentResponse{
		ID:        doc.ID().String(),
		FileName:  doc.FileName(),
		CreatedAt: doc.FormattedCreatedAt(),
	}
}
