package repository

import (
	"context"
	"errors"

	"github.com/yusufekoanggoro/oauth-app/backend/domain"
	"google.golang.org/api/idtoken"
)

type AuthRepository interface {
	VerifyGoogleToken(token string, clientID string) (*domain.GoogleUser, error)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) VerifyGoogleToken(token string, clientID string) (*domain.GoogleUser, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, token, clientID)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	user := &domain.GoogleUser{
		Email: payload.Claims["email"].(string),
		Name:  payload.Claims["name"].(string),
		Sub:   payload.Claims["sub"].(string),
	}

	return user, nil
}
