package log_service

import (
	log_domains "Log_service/internal/core/domains"
	log_core_logger "Log_service/internal/core/logger"
	"context"

	"github.com/google/uuid"
)

type LogService struct {
	cache    LogCache
	logInter Logger
}

//go:generate mockgen -source=service.go -destination=mocks/mock_logs_repo.go -package=logs_mocks
type LogCache interface {
	CheckUnique(
		ctx context.Context,
		uuid uuid.UUID,
	) (bool, error)
}

type Logger interface {
	LogMesage(event log_domains.LogEvent)
}

type LoggerRepository struct {
	log *log_core_logger.Logger
}

func NewLogService(
	cache LogCache,
	logInter Logger,
) *LogService {
	return &LogService{
		cache:    cache,
		logInter: logInter,
	}
}

func NewLoggerRepository(
	log *log_core_logger.Logger,
) *LoggerRepository {
	return &LoggerRepository{
		log: log,
	}
}
