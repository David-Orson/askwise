package presentation

import (
	"testing"
	"time"

	"askwise.com/m/v2/internal/document/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFromDomain_Table(t *testing.T) {
	now := time.Now()
	docID := uuid.New()
	projectID := uuid.New()
	userID := uuid.New()

	tests := []struct {
		name     string
		input    *domain.Document
		expected DocumentResponse
	}{
		{
			name: "valid document",
			input: domain.NewTestDocument(
				docID,
				projectID,
				userID,
				"report.pdf",
				now,
			),
			expected: DocumentResponse{
				ID:        docID.String(),
				FileName:  "report.pdf",
				CreatedAt: now.Format(time.RFC3339),
			},
		},
		{
			name:     "nil input",
			input:    nil,
			expected: DocumentResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := FromDomain(tt.input)
			assert.Equal(t, tt.expected, resp)
		})
	}
}
