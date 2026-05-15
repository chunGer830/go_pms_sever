package kk

import (
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	R *kafka.Reader
}

func GetReader(brokers []string, groupId, topic string) *KafkaReader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaReader{R: r}
}

func (r *KafkaReader) Close() {
	err := r.R.Close()
	if err != nil {
		return
	}
}
