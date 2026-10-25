package auth_postgres_repository

import (
	"context"
	"fmt"
)

func (r *AuthRepository) DispetchClassToTeacher(
	ctx context.Context,
	userID string,
	classID int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlUpdateRoomRequest := `
	UPDATE school.classes
	SET teacher_id = $1, version = version + 1
	WHERE id = $2;
	`

	if _, err := r.pool.Exec(ctxWithTime, sqlUpdateRoomRequest, userID, classID); err != nil {
		return fmt.Errorf("fail to dispetch class from database: %w", err)
	}

	return nil
}
