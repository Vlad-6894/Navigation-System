package auth_postgres_repository

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (r *AuthRepository) DispetchToParent(
	ctx context.Context,
	names auth_domains.Names,
	userID string,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlReq := `
	INSERT INTO school.parents (id, first_name, surname, middle_name)
	VALUES($1,$2,$3,$4);
	`

	if _, err := r.pool.Exec(ctxWithTime, sqlReq, userID, names.FirstName, names.Surname, names.MiddleName); err != nil {
		return fmt.Errorf("fail to dispetch to parent from database: %w", err)
	}

	return nil
}
