package auth_grpc_transport

import (
	"context"
	authv1 "navigation-proto/gen/go"
)

type AuthGrpcTransport struct {
	service AuthGrpcService
	authv1.UnimplementedAuthServiceServer
}

type AuthGrpcService interface {
	UpdateAccess(
		ctx context.Context,
		refreshToken string,
	) (string, error)

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
}

func NewAuthGrpcTransport(
	service AuthGrpcService,
) *AuthGrpcTransport {
	return &AuthGrpcTransport{
		service: service,
	}
}
