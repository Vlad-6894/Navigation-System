package auth_service

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (s *AuthService) RegisterProfile(
	ctx context.Context,
	profile auth_domains.Profile,
) (auth_domains.Profile, error) {
	if err := profile.Validate(); err != nil {
		return auth_domains.Profile{}, fmt.Errorf("fail validate profile: %w", err)
	}

	newProfile, err := s.db.RegisterProfile(ctx, profile)
	if err != nil {
		return auth_domains.Profile{}, fmt.Errorf("fail reg profile from repository: %w", err)
	}

	return newProfile, nil
}
