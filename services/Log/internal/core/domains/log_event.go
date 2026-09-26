package log_domains

import (
	log_core_logger "Log_service/internal/core/logger"

	"github.com/segmentio/kafka-go"
)

type LogEvent interface {
	With(log *log_core_logger.Logger) *log_core_logger.Logger
	GetMessage() kafka.Message
	GetErr() error
}
