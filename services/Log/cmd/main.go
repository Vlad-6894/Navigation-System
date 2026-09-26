package main

import (
	log_core_logger "Log_service/internal/core/logger"
	log_core_logger_config "Log_service/internal/core/logger/config"
	log_core_redis "Log_service/internal/core/repository/redis"
	log_core_kafka_transport "Log_service/internal/core/transport/kafka"
	log_core_kafka_transport_config "Log_service/internal/core/transport/kafka/config"
	log_redis_repository "Log_service/internal/features/repository/redis"
	log_service "Log_service/internal/features/service"
	log_kafka_transport "Log_service/internal/features/transport/kafka"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var (
	errorsChanSize = 3
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	wg := &sync.WaitGroup{}
	errorsChan := make(chan error, errorsChanSize)

	authLogger, err := log_core_logger.NewLogger(log_core_logger_config.NewAuthLoggerConfigMust())
	if err != nil {
		fmt.Println("init auth logger error!")
		os.Exit(1)
	}
	defer authLogger.Close()

	geoLogger, err := log_core_logger.NewLogger(log_core_logger_config.NewGeoLoggerConfigMust())
	if err != nil {
		fmt.Println("init geo logger error!")
		os.Exit(1)
	}
	defer geoLogger.Close()

	roomLogger, err := log_core_logger.NewLogger(log_core_logger_config.NewRoomLoggerConfigMust())
	if err != nil {
		fmt.Println("init room logger error!")
		os.Exit(1)
	}
	defer roomLogger.Close()

	client, err := log_core_redis.NewRedisClient(log_core_redis.NewRedisConfigMust())
	if err != nil {
		authLogger.Error("init redis client error!", zap.Error(err))
		os.Exit(1)
	}

	authRedisRepository := log_redis_repository.NewLogRedisRepository(client)
	authLoggerRepository := log_service.NewLoggerRepository(authLogger)

	geoRedisRepository := log_redis_repository.NewLogRedisRepository(client)
	geoLoggerRepository := log_service.NewLoggerRepository(geoLogger)

	roomRedisRepository := log_redis_repository.NewLogRedisRepository(client)
	roomLoggerRepository := log_service.NewLoggerRepository(roomLogger)

	authService := log_service.NewLogService(authRedisRepository, authLoggerRepository)

	geoService := log_service.NewLogService(geoRedisRepository, geoLoggerRepository)

	roomService := log_service.NewLogService(roomRedisRepository, roomLoggerRepository)

	authReaderConnfig := log_core_kafka_transport_config.NewAuthKafkaConsumerConfigMust()
	authReader := log_core_kafka_transport.NewLogKafkaReader(authReaderConnfig)
	defer authReader.Close()

	geoReaderConnfig := log_core_kafka_transport_config.NewGeoKafkaConsumerConfigMust()
	geoReader := log_core_kafka_transport.NewLogKafkaReader(geoReaderConnfig)
	defer geoLogger.Close()

	roomReaderConnfig := log_core_kafka_transport_config.NewRoomsKafkaConsumerConfigMust()
	roomReader := log_core_kafka_transport.NewLogKafkaReader(roomReaderConnfig)
	defer roomLogger.Close()

	authConsumer := log_kafka_transport.NewLogKafkaConsumer(authReader, authService, authLogger, wg)

	geoConsumer := log_kafka_transport.NewLogKafkaConsumer(geoReader, geoService, geoLogger, wg)

	roomConsumer := log_kafka_transport.NewLogKafkaConsumer(roomReader, roomService, roomLogger, wg)

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := authConsumer.Start(ctx, authReaderConnfig.Brokers, authReaderConnfig.Topic); err != nil {
			errorsChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := geoConsumer.Start(ctx, geoReaderConnfig.Brokers, geoReaderConnfig.Topic); err != nil {
			errorsChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := roomConsumer.Start(ctx, roomReaderConnfig.Brokers, roomReaderConnfig.Topic); err != nil {
			errorsChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		authLogger.Info("stop program by syscall")
	case err := <-errorsChan:
		authLogger.Error("stop program by error:", zap.Error(err))
		cancel()
	}

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		authLogger.Info("program stopped")
	case <-time.After(5 * time.Second):
		authLogger.Error("time is up!")
	}
}
