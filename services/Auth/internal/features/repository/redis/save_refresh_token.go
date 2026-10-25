package auth_redis_repository

import (
	auth_errors "Auth_service/internal/core/errors"
	"context"
	"encoding/json"
	"fmt"
	pkg_http_json "navigation-pkg/transport/http/json"
	"time"
)

func (r *AuthRedisRepository) SaveRefresToken(
	ctx context.Context,
	refreshToken string,
	userID string,
	role string,
	schoolID int,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.client.GetTimeout())
	defer cancel()

	tokenStruct := pkg_http_json.ProfileRefresh{
		UserID:   userID,
		Role:     role,
		SchoolID: schoolID,
	}

	b, err := json.Marshal(tokenStruct)
	if err != nil {
		return fmt.Errorf("fail to encode to json")
	}

	key := fmt.Sprintf("refresh:%s", refreshToken)

	success, err := r.client.SetNX(ctxWithTime, key, b, 7*24*time.Hour).Result()
	if err != nil {
		return fmt.Errorf("fail SetNx: %w", err)
	}

	if !success {
		return fmt.Errorf("fail SetNX: %w", auth_errors.ErrCreatedRefreshToken)
	}

	return nil
}
