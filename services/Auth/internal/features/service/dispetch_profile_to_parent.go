package auth_service

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (s *AuthService) DispetchToParent(
	ctx context.Context,
	names auth_domains.Names,
	userID string,
) error {
	if err := names.Validate(); err != nil {
		return fmt.Errorf("fail to validate names: %w", err)
	}

	if err := s.DispetchToParent(ctx, names, userID); err != nil {
		return fmt.Errorf("fail to dispetch to parent from database: %w", err)
	}

	return nil
}
