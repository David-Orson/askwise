package domain

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	validUUID := uuid.New()

	tests := []struct {
		name      string
		projectID uuid.UUID
		userID    uuid.UUID
		fileName  string
		wantErr   string
	}{
		{
			name:      "valid input (.pdf)",
			projectID: validUUID,
			userID:    validUUID,
			fileName:  "report.pdf",
			wantErr:   "",
		},
		{
			name:      "valid input (.md)",
			projectID: validUUID,
			userID:    validUUID,
			fileName:  "notes.md",
			wantErr:   "",
		},
		{
			name:      "empty projectID",
			projectID: uuid.Nil,
			userID:    validUUID,
			fileName:  "file.pdf",
			wantErr:   "projectID",
		},
		{
			name:      "empty userID",
			projectID: validUUID,
			userID:    uuid.Nil,
			fileName:  "file.pdf",
			wantErr:   "userID",
		},
		{
			name:      "empty fileName",
			projectID: validUUID,
			userID:    validUUID,
			fileName:  "",
			wantErr:   "fileName",
		},
		{
			name:      "whitespace fileName",
			projectID: validUUID,
			userID:    validUUID,
			fileName:  "   ",
			wantErr:   "fileName",
		},
		{
			name:      "invalid extension (.exe)",
			projectID: validUUID,
			userID:    validUUID,
			fileName:  "malware.exe",
			wantErr:   "supported extension",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := NewDocument(tt.projectID, tt.userID, tt.fileName)

			if tt.wantErr == "" {
				assert.NoError(t, err)
				assert.NotNil(t, doc)
			} else {
				assert.Error(t, err)
				assert.Nil(t, doc)
				assert.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestNewDocument_InvalidFileName(t *testing.T) {
	projectID := uuid.New()
	userID := uuid.New()

	tests := []struct {
		name      string
		fileName  string
		expectErr string
	}{
		{"empty", "", "fileName cannot be empty"},
		{"whitespace only", "   ", "fileName cannot be empty"},
		{"bad extension", "file.exe", "supported extension"},
		{"slash path", "foo/bar.pdf", "cannot contain slashes"},
		{"backslash path", "foo\\bar.pdf", "cannot contain slashes"},
		{"too long", strings.Repeat("a", 256) + ".pdf", "fileName is too long"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := NewDocument(projectID, userID, tt.fileName)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectErr)
			assert.Nil(t, doc)
		})
	}
}
