package auth_postgres_repository

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (r *AuthRepository) GetProfile(
	ctx context.Context,
	login string,
) (auth_domains.Profile, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlReq := `
		SELECT id, version, login, password, age, role, school_id FROM school.profiles
		WHERE login = $1;
		`

	row := r.pool.QueryRow(ctxWithTime, sqlReq, login)

	var profileModel ProfileModel

	if err := row.Scan(
		&profileModel.ID,
		&profileModel.Version,
		&profileModel.Login,
		&profileModel.Password,
		&profileModel.Age,
		&profileModel.Role,
		&profileModel.SchoolID,
	); err != nil {
		return auth_domains.Profile{}, fmt.Errorf("fail reg profile from database: %w", err)
	}

	profile := profileDomainFromModel(profileModel)

	return profile, nil
}
