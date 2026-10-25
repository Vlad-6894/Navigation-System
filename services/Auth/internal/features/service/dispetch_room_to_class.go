package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) DispetchRoomToClass(
	ctx context.Context,
	classID int,
	roomID int,
) error {
	if err := s.db.DispetchRoomToClass(ctx, classID, roomID); err != nil {
		return fmt.Errorf("fail to dispetch room from database: %w", err)
	}

	return nil
}
