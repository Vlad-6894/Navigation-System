//go:build integration

package log_redis_repository

import (
	log_core_redis "Log_service/internal/core/repository/redis"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	tc_redis "github.com/testcontainers/testcontainers-go/modules/redis"
)

var (
	keyPrefix         = "couriers:pool"
	redisImage        = "redis:8.6.1"
	testRedisPassword = "123"
	localHost         = "127.0.0.1"
)

func TestCheckUnique_integration(t *testing.T) {
	ctx := t.Context()

	redisContainer, err := tc_redis.Run(
		ctx,
		redisImage,
		testcontainers.WithCmd("redis-server", "--requirepass", testRedisPassword),
	)
	if err != nil {
		t.Fatalf("fail to up container: %v", err)
	}

	defer func() {
		if err := redisContainer.Terminate(context.Background()); err != nil {
			fmt.Println("fail close container", err)
		}
	}()

	port, err := redisContainer.MappedPort(t.Context(), "6379")
	if err != nil {
		t.Fatalf("fail to get port: %v", err)
	}

	os.Setenv("LOG_REDIS_HOST", localHost)
	os.Setenv("LOG_REDIS_PORT", port.Port())
	os.Setenv("LOG_REDIS_PASSWORD", testRedisPassword)
	os.Setenv("LOG_REDIS_TIMEOUT", "30s")

	defer func() {
		os.Unsetenv("LOG_REDIS_HOST")
		os.Unsetenv("LOG_REDIS_PORT")
		os.Unsetenv("LOG_REDIS_PASSWORD")
		os.Unsetenv("LOG_REDIS_TIMEOUT")
	}()

	client, err := log_core_redis.NewRedisClient(log_core_redis.NewRedisConfigMust())
	if err != nil {
		t.Fatalf("fail to get client: %v", err)
	}

	cache := NewLogRedisRepository(client)

	t.Run("test_check_unique_redis", func(t *testing.T) {
		uuid := uuid.New()

		ok, err := cache.CheckUnique(ctx, uuid)
		if err != nil {
			t.Errorf("fail check unique: %v", err)
		}

		if !ok {
			t.Errorf("fail check unique: wait ok=true, but returned ok=%v", ok)
		}

		ok, err = cache.CheckUnique(ctx, uuid)
		if err != nil {
			t.Errorf("fail check unique: %v", err)
		}

		if ok {
			t.Errorf("fail check unique: wait ok=false, but returned ok=%v", ok)
		}
	})
}
