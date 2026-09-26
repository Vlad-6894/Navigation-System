package log_kafka_transport

import (
	log_domains "Log_service/internal/core/domains"
	log_core_logger "Log_service/internal/core/logger"
	log_core_kafka_transport "Log_service/internal/core/transport/kafka"
	pkg_kafka_errors "Log_service/pkg/errors/kafka"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	maxLogsBatchSize = 100
	tickerTime       = 50 * time.Millisecond
	AuthServiceTopic = "log_auth"
	GeoServiceTopic  = "log_geo"
	RoomServiceTopic = "log_room"
	breakErr         = errors.New("break err")
	ununiqueErr      = errors.New("Ununique err")
)

type LogKafkaConsumer struct {
	*kafka.Reader
	service LogKafkaService
	log     *log_core_logger.Logger
	wg      *sync.WaitGroup
}

type LogKafkaService interface {
	CheckUnique(
		ctx context.Context,
		uuid uuid.UUID,
	) (bool, error)

	LogBatch(
		ctx context.Context,
		events []log_domains.LogEvent,
	) (int, error)
}

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
	c.log.Info("Start")
	num, err := log_core_kafka_transport.SearchPartitions(brokers, topic)
	if err != nil {
		return fmt.Errorf("fail get partitions num: %w", err)
	}

	if num == 0 {
		return fmt.Errorf("fail start consumer: %w", pkg_kafka_errors.NumPartitionsIsZero)
	}

	logChannel := make(chan log_domains.LogEvent, maxLogsBatchSize)
	ticker := time.NewTicker(tickerTime)

	for i := 0; i < num; i++ {
		c.wg.Add(1)
		go c.readPartition(ctx, topic, logChannel)
	}

	for {
		logBatch := make([]log_domains.LogEvent, 0, maxLogsBatchSize)

		select {
		case log, ok := <-logChannel:
			if !ok {
				c.log.Warn("close logs channel")
				return nil
			}

			if len(logBatch) == maxLogsBatchSize {
				commitNum, err := c.service.LogBatch(ctx, logBatch)
				if err != nil {
					c.log.Error("fail log batch", zap.Error(err))
				}

				if commitNum == -1 {
					continue
				}

				if err := c.Reader.CommitMessages(ctx, logBatch[commitNum].GetMessage()); err != nil {
					c.log.Error("fail commit batch", zap.Error(err))
				}

				continue
			}

			logBatch = append(logBatch, log)
		case <-ticker.C:
			commitNum, err := c.service.LogBatch(ctx, logBatch)
			if err != nil {
				c.log.Error("fail log batch", zap.Error(err))
			}

			if commitNum == -1 {
				continue
			}

			if err := c.Reader.CommitMessages(ctx, logBatch[commitNum].GetMessage()); err != nil {
				c.log.Error("fail commit batch", zap.Error(err))
			}

			continue

		case <-ctx.Done():
			c.log.Warn("context canceled")
			commitNum, err := c.service.LogBatch(ctx, logBatch)
			if err != nil {
				c.log.Error("fail log batch", zap.Error(err))
			}

			if commitNum == -1 {
				continue
			}

			if err := c.Reader.CommitMessages(ctx, logBatch[commitNum].GetMessage()); err != nil {
				c.log.Error("fail commit batch", zap.Error(err))
			}

			return nil
		}
	}

}

func (c *LogKafkaConsumer) readPartition(
	ctx context.Context,
	topic string,
	logChan chan log_domains.LogEvent,
) {
	defer c.wg.Done()

	for {
		if err := ctx.Err(); err != nil {
			c.log.Warn("stopped reader", zap.Error(err))
			return
		}

		msg, err := c.Reader.FetchMessage(ctx)
		if err != nil {
			c.log.Error("fail fetch message: ", zap.Error(err))
			continue
		}

		var event LogMessage

		switch topic {
		case AuthServiceTopic:
			Event, ok := event.(AuthLogMessage)
			if !ok {
				event = Event
			}
		}

		if err := json.Unmarshal(msg.Value, &event); err != nil {
			c.log.Error("fail umarshan message", zap.Error(err))
			msg := log_domains.NewFailMessage(event.GetUUID(), breakErr, msg)
			logChan <- msg
			continue
		}

		isUnique, err := c.service.CheckUnique(ctx, event.GetUUID())
		if err != nil {
			c.log.Error("fail check unique!")
			continue
		}

		if !isUnique {
			c.log.Warn("idemtece! Found a dublicate!")
			msg := log_domains.NewFailMessage(event.GetUUID(), ununiqueErr, msg)
			logChan <- msg

			continue
		}

		logEvent := event.LogEventFromMessage(msg)

		logChan <- logEvent
	}
}
