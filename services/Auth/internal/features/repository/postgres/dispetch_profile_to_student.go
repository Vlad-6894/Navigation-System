package auth_postgres_repository

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (r *AuthRepository) DispetchToStudent(
	ctx context.Context,
	names auth_domains.Names,
	userID string,
	classID int,
	parentID int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlReq := `
	INSERT INTO school.students (id, first_name, surname, middle_name, class_id, parent_id)
	VALUES($1,$2,$3,$4);
	`

	if _, err := r.pool.Exec(ctxWithTime, sqlReq, userID, names.FirstName, names.Surname, names.MiddleName, classID, parentID); err != nil {
		return fmt.Errorf("fail to dispetch to student from database: %w", err)
	}

	return nil
}
