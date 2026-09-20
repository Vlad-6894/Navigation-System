package log_core_kafka_transport_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type RoomsKafkaConsumerConfig struct {
	Brokers []string `envconfig:"ADDRESS" required:"true"`
	Topic   string   `envconfig:"ROOMS_TOPIC" required:"true"`
	GroupID string   `envconfig:"ROOMS_GROUP_ID" required:"true"`
}

func NewRoomsKafkaConsumerConfig() (RoomsKafkaConsumerConfig, error) {
	var config RoomsKafkaConsumerConfig

	if err := envconfig.Process("LOG_KAFKA_CONSUMER", &config); err != nil {
		return RoomsKafkaConsumerConfig{}, fmt.Errorf("fail process log rooms kafka consumer: %w", err)
	}

	return config, nil
}

func NewRoomsKafkaConsumerConfigMust() RoomsKafkaConsumerConfig {
	config, err := NewRoomsKafkaConsumerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func (c RoomsKafkaConsumerConfig) GetBrokers() []string {
	return c.Brokers
}

func (c RoomsKafkaConsumerConfig) GetTopic() string {
	return c.Topic
}

func (c RoomsKafkaConsumerConfig) GetGroupId() string {
	return c.GroupID
}
