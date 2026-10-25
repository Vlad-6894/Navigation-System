package auth_postgres_repository

import (
	"context"
	"fmt"
)

func (r *AuthRepository) UntieClassFromTeacher(
	ctx context.Context,
	classID int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlUpdateRoomRequest := `
	UPDATE school.classes
	SET teacher_id = NULL, version = version + 1
	WHERE id = $1;
	`

	if _, err := r.pool.Exec(ctxWithTime, sqlUpdateRoomRequest, classID); err != nil {
		return fmt.Errorf("fail to dispetch director from database: %w", err)
	}

	return nil
}
