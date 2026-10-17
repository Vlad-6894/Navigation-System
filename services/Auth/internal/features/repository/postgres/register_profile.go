package auth_postgres_repository

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	"fmt"
)

func (r *AuthRepository) RegisterProfile(
	ctx context.Context,
	profile auth_domains.Profile,
) (auth_domains.Profile, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	sqlReq := `
		INSERT INTO school.profiles(login, password, age, role, school_id)
		VALUES($1,$2,$3,$4,$5)
		RETURNING id, version, login, password, age, role, school_id 
		ON CONFLICT (login) DO NOTHING;
		`

	row := r.pool.QueryRow(ctxWithTime, sqlReq, profile.Login, profile.Password, profile.Age, profile.Role, profile.SchoolID)

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

	newProfile := profileDomainFromModel(profileModel)

	return newProfile, nil
}
