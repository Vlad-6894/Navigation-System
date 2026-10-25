package auth_grpc_transport

import (
	"context"
	authv1 "navigation-proto/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *AuthGrpcTransport) DispetchClassToTeacher(
	ctx context.Context,
	req *authv1.DispetchClassRequestMessage,
) (*authv1.ResponseMessage, error) {
	if err := t.service.DispetchClassToTeacher(ctx, req.UserId, int(req.ClassId)); err != nil {
		return &authv1.ResponseMessage{
			IsSuccess: false,
			Error:     err.Error(),
		}, status.Error(codes.Internal, "fail to dispetch class to teacher")
	}

	return &authv1.ResponseMessage{
		IsSuccess: true,
		Error:     "",
	}, nil
}
