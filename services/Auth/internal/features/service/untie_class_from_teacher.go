package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) UntieClassFromTeacher(
	ctx context.Context,
	classID int,
) error {
	if err := s.db.UntieClassFromTeacher(ctx, classID); err != nil {
		return fmt.Errorf("fail to dispetch class from database: %w", err)
	}

	return nil
}
