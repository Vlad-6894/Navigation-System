package auth_redis_repository

import pkg_redis "navigation-pkg/repository/redis"

type AuthRedisRepository struct {
	client pkg_redis.ClientCacheRedis
}

func NewAuthRedisRepository(
	client pkg_redis.ClientCacheRedis,
) *AuthRedisRepository {
	return &AuthRedisRepository{
		client: client,
	}
}
