package log_redis_repository

import log_core_redis "Log_service/internal/core/repository/redis"

type LogRedisRepository struct {
	client log_core_redis.ClientCacheRedis
}

func NewLogRedisRepository(
	client log_core_redis.ClientCacheRedis,
) *LogRedisRepository {
	return &LogRedisRepository{
		client: client,
	}
}
