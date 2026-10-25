package auth_postgres_repository

import (
	"context"
	"fmt"
)

func (r *AuthRepository) DispetchClassToStudent(
	ctx context.Context,
	userID string,
	classID int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlUpdateRoomRequest := `
	UPDATE school.students
	SET class_id = $1
	WHERE id = $2;
	`

	if _, err := r.pool.Exec(ctxWithTime, sqlUpdateRoomRequest, classID, userID); err != nil {
		return fmt.Errorf("fail to dispetch class from database: %w", err)
	}

	return nil
}
