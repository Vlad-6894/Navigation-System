package auth_postgres_repository

import (
	auth_domains "Auth_service/internal/core/domains"

	"github.com/google/uuid"
)

type ProfileModel struct {
	ID       uuid.UUID
	Version  int
	Login    string
	Password string
	Age      int
	Role     string
	SchoolID int
}

func profileDomainFromModel(model ProfileModel) auth_domains.Profile {
	profile := auth_domains.NewProfile(
		model.ID,
		model.Version,
		model.Login,
		model.Password,
		model.Age,
		model.Role,
		model.SchoolID,
	)

	return profile
}
