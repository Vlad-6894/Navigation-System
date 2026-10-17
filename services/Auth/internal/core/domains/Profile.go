package auth_domains

import (
	auth_errors "Auth_service/internal/core/errors"

	"github.com/google/uuid"
)

var (
	roles = []string{"ROOT", "DIRECTOR", "ADMIN", "TEACHER", "PARENT", "STUDENT"}
)

type Profile struct {
	ID       uuid.UUID
	Version  int
	Login    string
	Password string
	Age      int
	Role     string
	SchoolID int
}

func NewUninitilizedProfile(
	login string,
	password string,
	age int,
	role string,
	schoolID int,
) Profile {
	profile := Profile{
		ID:       uuid.New(),
		Version:  -1,
		Login:    login,
		Password: password,
		Age:      age,
		Role:     role,
		SchoolID: schoolID,
	}

	return profile
}

func NewProfile(
	id uuid.UUID,
	version int,
	login string,
	password string,
	age int,
	role string,
	schoolID int,
) Profile {
	profile := Profile{
		ID:       id,
		Version:  version,
		Login:    login,
		Password: password,
		Age:      age,
		Role:     role,
		SchoolID: schoolID,
	}

	return profile
}

func (p Profile) Validate() error {
	if len([]rune(p.Login)) < 8 || len([]rune(p.Login)) > 100 {
		return auth_errors.ErrInvalidLogin
	}

	if len([]rune(p.Password)) < 8 || len([]rune(p.Password)) > 100 {
		return auth_errors.ErrInvalidPassword
	}

	if p.Age < 0 {
		return auth_errors.ErrNegativeAge
	}

	if !p.containsRole() {
		return auth_errors.ErrNotFoundRole
	}

	return nil
}

func (p Profile) containsRole() bool {
	for _, v := range roles {
		if v == p.Role {
			return true
		}
	}

	return false
}
