package log_core_kafka_transport_config

type KafkaConsumerConfig interface {
	GetBrokers() []string
	GetTopic() string
	GetGroupId() string
}
