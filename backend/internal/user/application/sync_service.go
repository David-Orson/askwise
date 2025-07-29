package application

import (
	"context"
	"errors"

	"askwise.com/m/v2/internal/events"
	"askwise.com/m/v2/internal/user/domain"
	"askwise.com/m/v2/internal/user/ports"
	"gorm.io/gorm"
)

type UserService struct {
	Repo ports.UserRepository
	Bus  ports.EventBus
}

func NewUserService(repo ports.UserRepository, bus ports.EventBus) *UserService {
	return &UserService{Repo: repo, Bus: bus}
}

func (s *UserService) Sync(ctx context.Context, googleID, name, email, image string) (*domain.User, error) {
	user, err := s.Repo.FindByGoogleID(ctx, googleID)
	if err == nil {
		evt := events.UserSignedInEvent{
			UserID:   user.ID(),
			GoogleID: googleID,
			Name:     name,
			Email:    email,
			ImageURL: image,
		}
		s.Bus.Publish(ctx, events.EventUserSignedIn, evt)
		return user, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user = domain.NewUser(googleID, name, email, image)
	if err := s.Repo.Save(ctx, user); err != nil {
		return nil, err
	}

	evt := events.UserCreatedEvent{
		UserID:   user.ID(),
		GoogleID: googleID,
		Name:     name,
		Email:    email,
		ImageURL: image,
	}

	s.Bus.Publish(ctx, events.EventUserCreated, evt)
	return user, nil
}
