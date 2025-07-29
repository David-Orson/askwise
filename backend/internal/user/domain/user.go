package domain

import (
	"time"

	"askwise.com/m/v2/internal/shared/models"
	"github.com/google/uuid"
)

type User struct {
	Base     models.DomainBase
	googleID string
	name     string
	email    string
	imageURL string
}

func NewUser(googleID, name, email, imageURL string) *User {
	return &User{
		Base:     models.NewDomainBase(),
		googleID: googleID,
		name:     name,
		email:    email,
		imageURL: imageURL,
	}
}

func (u *User) ID() uuid.UUID {
	return u.Base.ID
}

func (u *User) GoogleID() string {
	return u.googleID
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) ImageURL() string {
	return u.imageURL
}

func (u *User) CreatedAt() time.Time {
	return u.Base.CreatedAt
}

func (u *User) UpdatedAt() time.Time {
	return u.Base.UpdatedAt
}

func (u *User) FormattedCreatedAt() string {
	return u.CreatedAt().Format(time.RFC3339)
}

func (u *User) FormattedUpdatedAt() string {
	return u.UpdatedAt().Format(time.RFC3339)
}

func ReconstructUser(
	id uuid.UUID,
	googleID, name, email, imageURL string,
	created, updated time.Time,
) *User {
	return &User{
		Base: models.DomainBase{
			ID:        id,
			CreatedAt: created,
			UpdatedAt: updated,
		},
		googleID: googleID,
		name:     name,
		email:    email,
		imageURL: imageURL,
	}
}

func NewTestUser(id uuid.UUID, googleID, name, email, imageURL string, created time.Time) *User {
	return &User{
		Base: models.DomainBase{
			ID:        id,
			CreatedAt: created,
			UpdatedAt: created,
		},
		googleID: googleID,
		name:     name,
		email:    email,
		imageURL: imageURL,
	}
}
