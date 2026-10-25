package auth_service

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (s *AuthService) DispetchToStudent(
	ctx context.Context,
	names auth_domains.Names,
	userID string,
	classID int,
	parentID int,
) error {
	if err := names.Validate(); err != nil {
		return fmt.Errorf("fail to validate names: %w", err)
	}

	if err := s.DispetchToStudent(ctx, names, userID, classID, parentID); err != nil {
		return fmt.Errorf("fail to dispetch to student from database: %w", err)
	}

	return nil
}
