package auth_grpc_transport

import (
	auth_errors "Auth_service/internal/core/errors"
	"context"
	"errors"
	authv1 "navigation-proto/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *AuthGrpcTransport) UpdateAccess(ctx context.Context, req *authv1.UpdateAccessRequestMessage) (*authv1.UpdateAccessResponseMessage, error) {
	newAccess, err := t.service.UpdateAccess(ctx, req.RefreshToken)
	if err != nil {
		if errors.Is(err, auth_errors.ErrNotFoundRefreshInCache) {
			return nil, status.Error(codes.NotFound, "resresh token ttl finished!")
		}
		return nil, status.Error(codes.Internal, "fail request")
	}

	return &authv1.UpdateAccessResponseMessage{
		AccessToken: newAccess,
	}, nil
}
