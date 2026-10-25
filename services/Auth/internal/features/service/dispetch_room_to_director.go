package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) DispetchRoomToDirector(
	ctx context.Context,
	userID string,
	roomID int,
) error {
	if err := s.db.DispetchRoomToDirector(ctx, userID, roomID); err != nil {
		return fmt.Errorf("fail to dispetch room from database: %w", err)
	}

	return nil
}
