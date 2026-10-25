package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) DispetchRoomToTeacher(
	ctx context.Context,
	userID string,
	roomID int,
) error {
	if err := s.db.DispetchRoomToTeacher(ctx, userID, roomID); err != nil {
		return fmt.Errorf("fail to dispetch room from database: %w", err)
	}

	return nil
}
