package kk

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type ToSendData struct {
	Topic string
	Key   []byte
	//json数据
	Data []byte
}

type KafkaWriter struct {
	w    *kafka.Writer
	data chan ToSendData
}

func GetWriter(addr string) *KafkaWriter {
	//TODO
	w := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	k := &KafkaWriter{
		w:    w,
		data: make(chan ToSendData, 100),
	}
	go k.sendKafka()
	return k
}

func (w *KafkaWriter) Send(data ToSendData) {
	w.data <- data
}

func (w *KafkaWriter) sendKafka() {
	for {
		select {
		case data := <-w.data:
			messages := []kafka.Message{
				{
					Topic: data.Topic,
					Key:   data.Key,
					Value: data.Data,
				},
			}

			var err error
			const retries = 3
			for i := 0; i < retries; i++ {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

				// attempt to create topic prior to publishing the message
				err = w.w.WriteMessages(ctx, messages...)
				cancel()

				if err == nil {
					break
				}

				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					continue
				}

				if err != nil {
					log.Printf("kafka send writeMsg err %s \n", err.Error())
				}
			}
		}
	}
}

func (w *KafkaWriter) Close() {
	if w.w != nil {
		w.w.Close()
	}
}
