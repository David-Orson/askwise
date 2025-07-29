package presentation

import (
	"askwise.com/m/v2/internal/user/domain"
)

type UserResponse struct {
	ID        string `json:"id"`
	GoogleID  string `json:"googleID"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImageURL  string `json:"imageURL"`
	CreatedAt string `json:"created"`
}

func FromDomain(user *domain.User) UserResponse {
	if user == nil {
		return UserResponse{}
	}

	return UserResponse{
		ID:        user.ID().String(),
		GoogleID:  user.GoogleID(),
		Name:      user.Name(),
		Email:     user.Email(),
		ImageURL:  user.ImageURL(),
		CreatedAt: user.FormattedCreatedAt(),
	}
}
