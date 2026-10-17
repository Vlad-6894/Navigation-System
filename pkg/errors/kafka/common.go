package pkg_kafka_errors

import "errors"

var NumPartitionsIsZero = errors.New("num partitions can not be 0!")
