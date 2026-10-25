package main

import (
	auth_postgres_repository "Auth_service/internal/features/repository/postgres"
	auth_redis_repository "Auth_service/internal/features/repository/redis"
	auth_service "Auth_service/internal/features/service"
	auth_grpc_transport "Auth_service/internal/features/transport/grpc"
	auth_http_transport "Auth_service/internal/features/transport/http"
	"context"
	"fmt"
	pkg_logger "navigation-pkg/logger"
	log_core_logger_config "navigation-pkg/logger/config"
	pkg_postgres_pool "navigation-pkg/repository/postgres/pool"
	pkg_redis "navigation-pkg/repository/redis"
	pkg_grpc_server "navigation-pkg/transport/grpc/server"
	pkg_http_middleware "navigation-pkg/transport/http/middleware"
	pkg_http_server "navigation-pkg/transport/http/server"
	authv1 "navigation-proto/gen/go"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	wg := &sync.WaitGroup{}

	logger, err := pkg_logger.NewLogger(log_core_logger_config.NewAuthLoggerConfigMust())
	if err != nil {
		fmt.Println("init logger error!")
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("init connection pool!")
	pool, err := pkg_postgres_pool.NewPostgresConnectionPool(ctx, pkg_postgres_pool.NewPostgresConnectionConfigMust())
	if err != nil {
		logger.Fatal("connection pool init error: ", zap.Error(err))
	}
	defer pool.Close()

	clientRedis, err := pkg_redis.NewRedisClient(pkg_redis.NewRedisConfigMust())
	if err != nil {
		logger.Fatal("fail init redis clent: ", zap.Error(err))
	}
	defer clientRedis.Close()

	authPostgresRepository := auth_postgres_repository.NewAuthRepository(pool)
	authRedisRepository := auth_redis_repository.NewAuthRedisRepository(clientRedis)

	authService := auth_service.NewAuthService(authPostgresRepository, authRedisRepository)

	authGrpcTransport := auth_grpc_transport.NewAuthGrpcTransport(authService)
	grpsServer := pkg_grpc_server.NewGRPCServer(pkg_grpc_server.NewHTTPServerConfigMust(), logger)
	authv1.RegisterAuthServiceServer(grpsServer.Server, authGrpcTransport)

	authHttpTransport := auth_http_transport.NewAuthHTTPTransport(authService)
	authHttpServer := pkg_http_server.NewHTTPServer(
		pkg_http_server.NewHTTPServerConfigMust(),
		logger,
		pkg_http_middleware.CORS(),
		pkg_http_middleware.RequestID(),
		pkg_http_middleware.LoggerForAuthService(logger),
		pkg_http_middleware.Trace(),
		pkg_http_middleware.Panic(),
	)

	apiVersionRouterStudents := pkg_http_server.NewAPIVersionRouter(pkg_http_server.ApiVersion1, pkg_http_server.Student)
	apiVersionRouterStudents.RegisterRoutes(authHttpTransport.Routes())

	done := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := authHttpServer.Run(ctx); err != nil {
			logger.Error("fail run http server: ", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := grpsServer.Run(ctx); err != nil {
			logger.Error("fail run http server: ", zap.Error(err))
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		logger.Info("program stopped")
	case <-time.After(5 * time.Second):
		logger.Error("time is up!")
	}
}
