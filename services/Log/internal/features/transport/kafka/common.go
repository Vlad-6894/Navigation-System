package log_kafka_transport

import (
	log_core_logger "Log_service/internal/core/logger"
	log_core_kafka_transport "Log_service/internal/core/transport/kafka"
	pkg_kafka_errors "Log_service/pkg/errors/kafka"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	maxLogsBatchSize = 100
	tickerTime       = 50 * time.Millisecond
)

type LogKafkaConsumer struct {
	*kafka.Reader
	service LogKafkaService
	log     *log_core_logger.Logger
	wg      *sync.WaitGroup
}

type LogKafkaService interface{}

func NewLogKafkaConsumer(
	reader *kafka.Reader,
	service LogKafkaService,
	log *log_core_logger.Logger,
	wg *sync.WaitGroup,
) *LogKafkaConsumer {
	return &LogKafkaConsumer{
		Reader:  reader,
		service: service,
		log:     log,
		wg:      wg,
	}
}

func (c *LogKafkaConsumer) Start(
	ctx context.Context,
	brokers []string,
	topic string,
) error {
	num, err := log_core_kafka_transport.SearchPartitions(brokers, topic)
	if err != nil {
		return fmt.Errorf("fail get partitions num: %w", err)
	}

	if num == 0 {
		return fmt.Errorf("fail start consumer: %w", pkg_kafka_errors.NumPartitionsIsZero)
	}

	logChannel := make(chan LogEvent, maxLogsBatchSize)
	ticker := time.NewTicker(tickerTime)
}
