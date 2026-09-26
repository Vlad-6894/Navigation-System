package log_domains

import (
	log_core_logger "Log_service/internal/core/logger"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type FailMessage struct {
	UUID    uuid.UUID
	Err     error
	Message kafka.Message
}

func NewFailMessage(
	uuid uuid.UUID,
	err error,
	message kafka.Message,
) FailMessage {
	event := FailMessage{
		UUID:    uuid,
		Err:     err,
		Message: message,
	}

	return event
}

func (e FailMessage) With(log *log_core_logger.Logger) *log_core_logger.Logger {
	logger := log.With()

	return logger
}

func (e FailMessage) GetMessage() kafka.Message {
	return e.Message
}

func (e FailMessage) GetErr() error {
	return e.Err
}
