package log_kafka_transport

import (
	log_domains "Log_service/internal/core/domains"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type AuthLogMessage struct {
	UUID      uuid.UUID `json:"uuid"`
	RequestID string    `json:"request_id"`
	URL       string    `json:"url"`
	Time      time.Time `json:"time_now"`
}

type LogMessage interface {
	GetUUID() uuid.UUID
	LogEventFromMessage(messageg kafka.Message) log_domains.LogEvent
}

func (msg AuthLogMessage) LogEventFromMessage(message kafka.Message) log_domains.LogEvent {
	event := log_domains.NewAuthLogEvent(msg.UUID, msg.RequestID, msg.URL, nil, message)
	return event
}

func (m AuthLogMessage) GetUUID() uuid.UUID {
	return m.UUID
}
