package adapters

import (
	"context"
	"testing"
	"time"

	"askwise.com/m/v2/internal/document/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&DocumentRecord{})
	assert.NoError(t, err)

	return db
}

func TestPostgresDocumentRepository_SaveAndFindByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostgresDocumentRepository(db)

	ctx := context.Background()

	doc := domain.NewTestDocument(
		uuid.New(),
		uuid.New(),
		uuid.New(),
		"test.pdf",
		time.Now(),
	)

	err := repo.Save(ctx, doc)
	assert.NoError(t, err)

	found, err := repo.FindByID(ctx, doc.ID())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, doc.ID(), found.ID())
	assert.Equal(t, doc.FileName(), found.FileName())
}
