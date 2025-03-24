package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yusufekoanggoro/oauth-app/backend/config"
	"github.com/yusufekoanggoro/oauth-app/backend/repository"
)

type AuthUsecase interface {
	Authenticate(token string) (string, error)
}

type authUsecase struct {
	repo repository.AuthRepository
	cfg  *config.Config
}

func NewAuthUsecase(repo repository.AuthRepository, cfg *config.Config) AuthUsecase {
	return &authUsecase{repo: repo, cfg: cfg}
}

func (u *authUsecase) Authenticate(token string) (string, error) {
	user, err := u.repo.VerifyGoogleToken(token, u.cfg.ClientID)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.Sub,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(u.cfg.SecretKey))
	if err != nil {
		return "", errors.New("failed to create JWT")
	}

	return tokenString, nil
}
