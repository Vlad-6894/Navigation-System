package log_core_kafka_transport_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type GeoKafkaConsumerConfig struct {
	Brokers []string `envconfig:"ADDRESS" required:"true"`
	Topic   string   `envconfig:"GEO_TOPIC" required:"true"`
	GroupID string   `envconfig:"GEO_GROUP_ID" required:"true"`
}

func NewGeoKafkaConsumerConfig() (GeoKafkaConsumerConfig, error) {
	var config GeoKafkaConsumerConfig

	if err := envconfig.Process("LOG_KAFKA_CONSUMER", &config); err != nil {
		return GeoKafkaConsumerConfig{}, fmt.Errorf("fail process log geo kafka consumer: %w", err)
	}

	return config, nil
}

func NewGeoKafkaConsumerConfigMust() GeoKafkaConsumerConfig {
	config, err := NewGeoKafkaConsumerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func (c GeoKafkaConsumerConfig) GetBrokers() []string {
	return c.Brokers
}

func (c GeoKafkaConsumerConfig) GetTopic() string {
	return c.Topic
}

func (c GeoKafkaConsumerConfig) GetGroupId() string {
	return c.GroupID
}
