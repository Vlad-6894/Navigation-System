package auth_service

import (
	"context"
	"fmt"
	pkg_jwt "navigation-pkg/jwt"
)

func (s *AuthService) UpdateAccess(
	ctx context.Context,
	refreshToken string,
) (string, error) {
	userID, role, schoolID, err := s.cache.FindRefrshToken(ctx, refreshToken)
	if err != nil {
		return "", fmt.Errorf("fail to find refresh token: %w", err)
	}

	accessToken, err := pkg_jwt.GenerateAccessToken(userID, role, schoolID)
	if err != nil {
		return "", fmt.Errorf("fail generate access token: %w", err)
	}

	return accessToken, nil
}
