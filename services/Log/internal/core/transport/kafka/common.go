package log_core_kafka_transport

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

var (
	networkProto = "tcp"
)

func SearchPartitions(address []string, topic string) (int, error) {
	conn, err := kafka.Dial(networkProto, address[0])
	if err != nil {
		return 0, fmt.Errorf("fail conn to broker: %w", err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions(topic)
	if err != nil {
		return 0, fmt.Errorf("fail read partitions from topic: %w", err)
	}

	num := len(partitions)

	return num, nil
}
