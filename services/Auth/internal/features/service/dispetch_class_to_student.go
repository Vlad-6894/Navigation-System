package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) DispetchClassToStudent(
	ctx context.Context,
	userID string,
	classID int,
) error {
	if err := s.db.DispetchClassToStudent(ctx, userID, classID); err != nil {
		return fmt.Errorf("fail to dispetch class from database: %w", err)
	}

	return nil
}
