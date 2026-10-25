package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) DispetchClassToTeacher(
	ctx context.Context,
	userID string,
	classID int,
) error {
	if err := s.db.DispetchClassToTeacher(ctx, userID, classID); err != nil {
		return fmt.Errorf("fail to dispetch class from database: %w", err)
	}

	return nil
}
