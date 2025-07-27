package application

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUploadDocument(t *testing.T) {
	validProjectID := uuid.New()
	validUserID := uuid.New()

	tests := []struct {
		name        string
		projectID   uuid.UUID
		userID      uuid.UUID
		fileName    string
		mockSaveErr error
		mockPubErr  error
		expectErr   bool
		expectDoc   bool
	}{
		{
			name:        "success",
			projectID:   validProjectID,
			userID:      validUserID,
			fileName:    "file.pdf",
			mockSaveErr: nil,
			mockPubErr:  nil,
			expectErr:   false,
			expectDoc:   true,
		},
		{
			name:        "repo fails to save",
			projectID:   validProjectID,
			userID:      validUserID,
			fileName:    "file.pdf",
			mockSaveErr: errors.New("db error"),
			expectErr:   true,
			expectDoc:   false,
		},
		{
			name:        "bus fails to publish",
			projectID:   validProjectID,
			userID:      validUserID,
			fileName:    "file.pdf",
			mockSaveErr: nil,
			mockPubErr:  errors.New("redis error"),
			expectErr:   true,
			expectDoc:   false,
		},
		{
			name:      "invalid (empty filename)",
			projectID: validProjectID,
			userID:    validUserID,
			fileName:  "",
			expectErr: true,
			expectDoc: false,
		},
		{
			name:      "invalid projectID (uuid.Nil)",
			projectID: uuid.Nil,
			userID:    validUserID,
			fileName:  "file.pdf",
			expectErr: true,
			expectDoc: false,
		},
		{
			name:      "invalid userID (uuid.Nil)",
			projectID: validProjectID,
			userID:    uuid.Nil,
			fileName:  "file.pdf",
			expectErr: true,
			expectDoc: false,
		},
		{
			name:      "invalid file extension",
			projectID: validProjectID,
			userID:    validUserID,
			fileName:  "malware.exe",
			expectErr: true,
			expectDoc: false,
		},
		{
			name:      "no extension",
			projectID: validProjectID,
			userID:    validUserID,
			fileName:  "README",
			expectErr: true,
			expectDoc: false,
		},
		{
			name:      "whitespace-only filename",
			projectID: validProjectID,
			userID:    validUserID,
			fileName:  "    ",
			expectErr: true,
			expectDoc: false,
		},
		{
			name:        "uppercase extension (.PDF)",
			projectID:   validProjectID,
			userID:      validUserID,
			fileName:    "UPPER.PDF",
			mockSaveErr: nil,
			mockPubErr:  nil,
			expectErr:   false,
			expectDoc:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(MockRepo)
			bus := new(MockBus)
			svc := NewDocumentService(repo, bus)

			if tt.mockSaveErr == nil && tt.fileName != "" {
				repo.On("Save", mock.Anything, mock.Anything).Return(nil)
			} else if tt.fileName != "" {
				repo.On("Save", mock.Anything, mock.Anything).Return(tt.mockSaveErr)
			}

			if tt.mockPubErr == nil && tt.mockSaveErr == nil && tt.fileName != "" {
				bus.On("Publish", mock.Anything, "document.uploaded", mock.Anything).Return(nil)
			} else if tt.mockPubErr != nil && tt.fileName != "" && tt.mockSaveErr == nil {
				bus.On("Publish", mock.Anything, "document.uploaded", mock.Anything).Return(tt.mockPubErr)
			}

			doc, err := svc.UploadDocument(context.Background(), tt.projectID, tt.userID, tt.fileName)

			if tt.expectErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "expected no error but got one")
			}

			if tt.expectDoc {
				assert.NotNil(t, doc, "expected document but got nil")
				if doc != nil {
					assert.Equal(t, tt.fileName, doc.FileName())
				}
			} else {
				assert.Nil(t, doc, "expected nil document but got one")
			}
		})
	}
}
