package auth_grpc_transport

import (
	"context"
	authv1 "navigation-proto/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *AuthGrpcTransport) DispetchRoomToAdmin(
	ctx context.Context,
	req *authv1.DispetchRoomRequestMessage,
) (*authv1.ResponseMessage, error) {
	if err := t.service.DispetchRoomToAdmin(ctx, req.UserId, int(req.RoomId)); err != nil {
		return &authv1.ResponseMessage{
			IsSuccess: false,
			Error:     err.Error(),
		}, status.Error(codes.Internal, "fail to dispetch room to admin")
	}

	return &authv1.ResponseMessage{
		IsSuccess: true,
		Error:     "",
	}, nil
}
