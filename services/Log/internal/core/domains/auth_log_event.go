package log_domains

import (
	log_core_logger "Log_service/internal/core/logger"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type AuthLogEvent struct {
	UUID      uuid.UUID
	RequestID string
	URL       string
	Err       error
	Message   kafka.Message
}

func NewAuthLogEvent(
	uuid uuid.UUID,
	requestID string,
	url string,
	err error,
	message kafka.Message,
) AuthLogEvent {
	event := AuthLogEvent{
		UUID:      uuid,
		RequestID: requestID,
		URL:       url,
		Err:       err,
		Message:   message,
	}

	return event
}

func (e AuthLogEvent) With(log *log_core_logger.Logger) *log_core_logger.Logger {
	logger := log.With(
		zap.Any("uuid", e.UUID),
		zap.String("request_id", e.RequestID),
		zap.String("url", e.URL),
	)

	return logger
}

func (e AuthLogEvent) GetMessage() kafka.Message {
	return e.Message
}

func (e AuthLogEvent) GetErr() error {
	return e.Err
}
