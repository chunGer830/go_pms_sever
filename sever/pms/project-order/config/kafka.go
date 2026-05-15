package config

import (
	"context"
	"go.uber.org/zap"
	"pms.com/project-common/kk"
)

var KW *kk.KafkaWriter

func InitKafkaWriter() func() {
	KW := kk.GetWriter("localhost:9092")
	return KW.Close
}

func SendLog(data []byte) {
	KW.Send(kk.LogData{
		Topic: "pms_log",
		Data:  data,
	})
}

type KafkaCache struct {
	R *kk.KafkaReader
}

func (c *KafkaCache) DeleteCache() {
	for {
		message, err := c.R.R.ReadMessage(context.Background())
		if err != nil {
			zap.L().Error("DeleteCache error", zap.Error(err))
			continue
		}
		if "task" == string(message.Value) {

		}
	}
}

func NewCacheReader() *KafkaCache {
	reader := kk.GetReader([]string{"localhost:9092"}, "cache_group", "pms_cache")
	return &KafkaCache{
		R: reader,
	}
}
