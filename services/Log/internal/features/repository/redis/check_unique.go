package log_redis_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	TTL               = 5 * time.Minute
	prefixIdempotensy = "idempotency:log"
)

func (r *LogRedisRepository) CheckUnique(
	ctx context.Context,
	uuid uuid.UUID,
) (bool, error) {
	ctxWithTime, cacel := context.WithTimeout(ctx, r.client.GetTimeout())
	defer cacel()

	key := fmt.Sprintf("%s:%s", prefixIdempotensy, uuid.String())

	ok, err := r.client.SetNX(ctxWithTime, key, "1", TTL).Result()
	if err != nil {
		return false, fmt.Errorf("fail SetNX: %w", err)
	}

	return ok, err
}
