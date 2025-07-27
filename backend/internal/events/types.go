package events

import "github.com/google/uuid"

const (
	EventUserCreated      = "user.created"
	EventUserSignedIn     = "user.signed_in"
	EventProjectCreated   = "project.created"
	EventProjectUpdated   = "project.updated"
	EventProjectDeleted   = "project.deleted"
	EventDocumentUploaded = "document.uploaded"
	EventDocumentChunked  = "document.chunked"
	EventDocumentDeleted  = "document.deleted"
	EventQuestionAsked    = "question.asked"
	EventAnswerGenerated  = "answer.generated"
)

type DocumentUploadedEvent struct {
	DocumentID uuid.UUID `json:"document_id"`
	ProjectID  uuid.UUID `json:"project_id"`
	UserID     uuid.UUID `json:"user_id"`
	FileName   string    `json:"file_name"`
	UploadedAt string    `json:"uploaded_at"`
}
