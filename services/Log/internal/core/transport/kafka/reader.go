package log_core_kafka_transport

import (
	log_core_kafka_transport_config "Log_service/internal/core/transport/kafka/config"

	"github.com/segmentio/kafka-go"
)

func NewLogKafkaReader(config log_core_kafka_transport_config.KafkaConsumerConfig) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        config.GetBrokers(),
		Topic:          config.GetTopic(),
		GroupID:        config.GetGroupId(),
		CommitInterval: 0,
	})

	return reader
}
