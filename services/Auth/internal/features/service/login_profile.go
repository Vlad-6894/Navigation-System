package auth_service

import (
	auth_errors "Auth_service/internal/core/errors"
	"context"
	"fmt"
	pkg_jwt "navigation-pkg/jwt"
)

func (s *AuthService) LoginProfile(
	ctx context.Context,
	login string,
	password string,
) (pkg_jwt.TokensPair, error) {
	profile, err := s.db.GetProfile(ctx, login)
	if err != nil {
		return pkg_jwt.TokensPair{}, fmt.Errorf("fail login profile: %w", err)
	}

	if profile.Password != password {
		return pkg_jwt.TokensPair{}, auth_errors.ErrInvalidPassword
	}

	tokenPair, err, accessClaims := pkg_jwt.GenerateTokensPare(profile.ID.String(), profile.Role, profile.SchoolID)
	if err != nil {
		return pkg_jwt.TokensPair{}, fmt.Errorf("fail to get token pair: %w", err)
	}

	if err := s.cache.SaveRefresToken(ctx, tokenPair.RefreshToken, accessClaims.UserID, accessClaims.Role, accessClaims.SchoolID); err != nil {
		return pkg_jwt.TokensPair{}, fmt.Errorf("fail to save token pair to redis: %w", err)
	}

	return tokenPair, nil
}
