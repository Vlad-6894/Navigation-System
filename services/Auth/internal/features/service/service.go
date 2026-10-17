package auth_service

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
)

type AuthService struct {
	db AuthDatabase
}

type AuthDatabase interface {
	RegisterProfile(
		ctx context.Context,
		profile auth_domains.Profile,
	) (auth_domains.Profile, error)
}

func NewAuthService(db AuthDatabase) *AuthService {
	return &AuthService{
		db: db,
	}
}
