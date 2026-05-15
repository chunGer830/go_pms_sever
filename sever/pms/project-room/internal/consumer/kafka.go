package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"pms.com/project-common/kk"
	"pms.com/project-grpc/order/order_inf"
	"pms.com/project-room/internal/dao"
	"pms.com/project-room/internal/data/room_type_data"
	"time"
)

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw = kk.GetWriter("localhost:9092")
	return kw.Close
}

func SendLog(data []byte) {
	kw.Send(kk.LogData{
		Topic: "pms_log",
		Data:  data,
	})
}

type KafkaCache struct {
	R *kk.KafkaReader
}

func NewCacheReader() *KafkaCache {
	reader := kk.GetReader([]string{"localhost:9092"}, "cache_group", "pms_cache")
	return &KafkaCache{
		R: reader,
	}
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

func SendOrderServiceMsg(data []byte) {
	kw.Send(kk.LogData{
		Topic: "orderService",
		Data:  data,
	})
}

func NewOrderServiceReader() *KafkaCache {
	reader := kk.GetReader([]string{"localhost:9092"}, "orderService_group", "orderService")
	return &KafkaCache{
		R: reader,
	}
}

func (c *KafkaCache) OrderService() {
	for {
		message, err := c.R.R.ReadMessage(context.Background())
		fmt.Println("message.Key:", string(message.Key))
		if err != nil {
			zap.L().Error("OrderService error", zap.Error(err))
			continue
		}

		if "Service" == string(message.Key) {

			var data room_type_data.RoomGuestStay
			if err := json.Unmarshal(message.Value, &data); err != nil {
				zap.L().Error("unmarshal message error", zap.Error(err))
				continue
			}
			orderServiceClient := dao.OrderServiceClient
			msg := &order_inf.OrderInfMessage{
				HotelId:      data.HotelID,
				GuestRoomNo:  data.GuestRoomNo,
				GuestName:    data.GuestName,
				GuestIdNo:    data.GuestIDNo,
				RealPrice:    data.RealPrice,
				Mobile:       data.Mobile,
				CheckInTime:  data.CheckInTime,
				CheckOutTime: data.CheckOutTime,
			}
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

			_, err = orderServiceClient.OrderInf(ctx, msg)
			cancel()
			if err != nil {
				zap.L().Error(" err ", zap.Error(err))
			}
		}

	}
}
