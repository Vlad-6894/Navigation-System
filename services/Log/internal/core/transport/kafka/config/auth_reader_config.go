package log_core_kafka_transport_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AuthKafkaConsumerConfig struct {
	Brokers []string `envconfig:"ADDRESS" required:"true"`
	Topic   string   `envconfig:"AUTH_TOPIC" required:"true"`
	GroupID string   `envconfig:"AUTH_GROUP_ID" required:"true"`
}

func NewAuthKafkaConsumerConfig() (AuthKafkaConsumerConfig, error) {
	var config AuthKafkaConsumerConfig

	if err := envconfig.Process("LOG_KAFKA_CONSUMER", &config); err != nil {
		return AuthKafkaConsumerConfig{}, fmt.Errorf("fail process log auth kafka consumer: %w", err)
	}

	return config, nil
}

func NewAuthKafkaConsumerConfigMust() AuthKafkaConsumerConfig {
	config, err := NewAuthKafkaConsumerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func (c AuthKafkaConsumerConfig) GetBrokers() []string {
	return c.Brokers
}

func (c AuthKafkaConsumerConfig) GetTopic() string {
	return c.Topic
}

func (c AuthKafkaConsumerConfig) GetGroupId() string {
	return c.GroupID
}
