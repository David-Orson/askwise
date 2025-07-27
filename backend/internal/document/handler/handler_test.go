package handler

import (
	"mime/multipart"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"askwise.com/m/v2/internal/document/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUploadHandler(t *testing.T) {
	type testCase struct {
		name         string
		setupLocals  func(*fiber.Ctx)
		projectID    string
		fileName     string
		mockService  func(*MockDocService)
		expectStatus int
	}

	userID := uuid.New()
	projectID := uuid.New()

	tests := []testCase{
		{
			name:      "success",
			projectID: projectID.String(),
			fileName:  "test.pdf",
			setupLocals: func(c *fiber.Ctx) {
				c.Locals("userID", userID.String())
			},
			mockService: func(svc *MockDocService) {
				doc := domain.NewTestDocument(uuid.New(), uuid.MustParse(projectID.String()), userID, "test.pdf", time.Now())
				svc.On("UploadDocument", mock.Anything, uuid.MustParse(projectID.String()), userID, "test.pdf").
					Return(doc, nil)
			},
			expectStatus: 200,
		},
		{
			name:         "unauthenticated",
			projectID:    projectID.String(),
			fileName:     "test.pdf",
			setupLocals:  func(c *fiber.Ctx) {},
			mockService:  func(svc *MockDocService) {},
			expectStatus: 401,
		},
		{
			name:      "invalid_userID",
			projectID: projectID.String(),
			fileName:  "test.pdf",
			setupLocals: func(c *fiber.Ctx) {
				c.Locals("userID", "not-a-uuid")
			},
			mockService:  func(svc *MockDocService) {},
			expectStatus: 400,
		},
		{
			name:         "invalid_projectID",
			projectID:    "not-a-uuid",
			fileName:     "test.pdf",
			setupLocals:  func(c *fiber.Ctx) { c.Locals("userID", userID.String()) },
			mockService:  func(svc *MockDocService) {},
			expectStatus: 400,
		},
		{
			name:         "missing_file",
			projectID:    projectID.String(),
			fileName:     "", // simulate no file
			setupLocals:  func(c *fiber.Ctx) { c.Locals("userID", userID.String()) },
			mockService:  func(svc *MockDocService) {},
			expectStatus: 400,
		},
		{
			name:      "service_error",
			projectID: projectID.String(),
			fileName:  "fail.pdf",
			setupLocals: func(c *fiber.Ctx) {
				c.Locals("userID", userID.String())
			},
			mockService: func(svc *MockDocService) {
				svc.On("UploadDocument", mock.Anything, uuid.MustParse(projectID.String()), userID, "fail.pdf").
					Return(nil, assert.AnError)
			},
			expectStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			svc := new(MockDocService)
			handler := NewDocumentHandler(svc)

			tt.mockService(svc)

			app.Use(func(c *fiber.Ctx) error {
				tt.setupLocals(c)
				return c.Next()
			})
			app.Post("/api/projects/:projectID/upload", handler.Upload)

			var body *strings.Reader
			var contentType string

			if tt.fileName != "" {
				body, contentType = createMultipartForm(t, tt.fileName, "file")
			} else {
				body = strings.NewReader("")
				contentType = "multipart/form-data"
			}

			req := httptest.NewRequest("POST", "/api/projects/"+tt.projectID+"/upload", body)
			req.Header.Set("Content-Type", contentType)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectStatus, resp.StatusCode)
		})
	}
}

func createMultipartForm(t *testing.T, fileName, fieldName string) (*strings.Reader, string) {
	var b strings.Builder
	writer := multipart.NewWriter(&b)

	part, err := writer.CreateFormFile(fieldName, fileName)
	assert.NoError(t, err)
	_, _ = part.Write([]byte("fake content"))

	assert.NoError(t, writer.Close())

	return strings.NewReader(b.String()), writer.FormDataContentType()
}
