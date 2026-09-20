package log_kafka_transport

import log_core_logger "Log_service/internal/core/logger"

type LogEvent interface {
	With(*log_core_logger.Logger) *log_core_logger.Logger
}
