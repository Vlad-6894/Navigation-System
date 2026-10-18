package auth_service

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
)

type AuthService struct {
	db    AuthDatabase
	cache AuthCache
}

type AuthDatabase interface {
	RegisterProfile(
		ctx context.Context,
		profile auth_domains.Profile,
	) (auth_domains.Profile, error)

	GetProfile(
		ctx context.Context,
		login string,
	) (auth_domains.Profile, error)
}

type AuthCache interface {
	SaveRefresToken(
		ctx context.Context,
		refreshToken string,
		userID string,
		role string,
		schoolID int,
	) error
}

func NewAuthService(db AuthDatabase, cache AuthCache) *AuthService {
	return &AuthService{
		db:    db,
		cache: cache,
	}
}
