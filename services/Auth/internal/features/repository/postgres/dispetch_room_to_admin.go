package auth_postgres_repository

import (
	"context"
	"fmt"
)

func (r *AuthRepository) DispetchRoomToAdmin(
	ctx context.Context,
	userID string,
	roomID int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlUpdateRoomRequest := `
	UPDATE school.admins
	SET room_id = $1, version = version + 1
	WHERE id = $2;
	`

	if _, err := r.pool.Exec(ctxWithTime, sqlUpdateRoomRequest, roomID, userID); err != nil {
		return fmt.Errorf("fail to dispetch admin from database: %w", err)
	}

	return nil
}
