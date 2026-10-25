package auth_redis_repository

import (
	auth_errors "Auth_service/internal/core/errors"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pkg_http_json "navigation-pkg/transport/http/json"

	"github.com/redis/go-redis/v9"
)

func (r *AuthRedisRepository) FindRefrshToken(
	ctx context.Context,
	refreshToken string,
) (string, string, int, error) {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.client.GetTimeout())
	defer cancel()

	key := fmt.Sprintf("refresh:%s", refreshToken)

	res, err := r.client.Get(ctxWithTime, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", "", 0, fmt.Errorf("fail to get from cache: %v : %w", err, auth_errors.ErrNotFoundRefreshInCache)
		}

		return "", "", 0, err
	}

	var refreshStruct pkg_http_json.ProfileRefresh

	if err := json.Unmarshal([]byte(res), &refreshStruct); err != nil {
		return "", "", 0, fmt.Errorf("fail unmarshal resonse from redis: %w", err)
	}

	return refreshStruct.UserID, refreshStruct.Role, refreshStruct.SchoolID, nil
}
