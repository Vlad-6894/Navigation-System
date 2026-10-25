package auth_service

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
)

type AuthService struct {
	db    AuthDatabase
	cache AuthCache
}

type AuthDatabase interface {
	RegisterProfile(
		ctx context.Context,
		profile auth_domains.Profile,
	) (auth_domains.Profile, error)

	GetProfile(
		ctx context.Context,
		login string,
	) (auth_domains.Profile, error)

	DispetchRoomToDirector(
		ctx context.Context,
		userID string,
		roomID int,
	) error

	DispetchRoomToAdmin(
		ctx context.Context,
		userID string,
		roomID int,
	) error

	DispetchRoomToTeacher(
		ctx context.Context,
		userID string,
		roomID int,
	) error

	DispetchClassToTeacher(
		ctx context.Context,
		userID string,
		classID int,
	) error

	DispetchRoomToClass(
		ctx context.Context,
		classID int,
		roomID int,
	) error

	DispetchClassToStudent(
		ctx context.Context,
		userID string,
		classID int,
	) error

	UntieClassFromTeacher(
		ctx context.Context,
		classID int,
	) error

	DispetchToParent(
		ctx context.Context,
		names auth_domains.Names,
		userID string,
	) error

	DispetchToStudent(
		ctx context.Context,
		names auth_domains.Names,
		userID string,
		classID int,
		parentID int,
	) error
}

type AuthCache interface {
	SaveRefresToken(
		ctx context.Context,
		refreshToken string,
		userID string,
		role string,
		schoolID int,
	) error

	FindRefrshToken(
		ctx context.Context,
		refreshToken string,
	) (string, string, int, error)
}

func NewAuthService(db AuthDatabase, cache AuthCache) *AuthService {
	return &AuthService{
		db:    db,
		cache: cache,
	}
}
