package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"pms.com/project-common/kk"
	"pms.com/project-grpc/order/order_inf"
	"pms.com/project-room/config"
	"pms.com/project-room/internal/dao"
	"pms.com/project-room/internal/data/room_type_data"
	"time"
)

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw = kk.GetWriter(config.C.KafkaConfig.Addr)
	return kw.Close
}

type KafkaCache struct {
	R *kk.KafkaReader
}

func SendOrderServiceMsg(data []byte) {
	kw.Send(kk.ToSendData{
		Topic: "orderService",
		Key:   []byte("Service"),
		Data:  data,
	})
}

func NewOrderServiceReader() *KafkaCache {
	//TODO  修改config
	reader := kk.GetReader([]string{config.C.KafkaConfig.Addr}, "orderService_group", "orderService")
	return &KafkaCache{
		R: reader,
	}
}

func (c *KafkaCache) OrderService() {
	for {
		// 1. 读取消息（超时只表示无新消息）
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		message, err := c.R.R.ReadMessage(ctx)
		cancel()

		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				// 超时无消息，正常继续
				continue
			}
			zap.L().Error("OrderService error", zap.Error(err))
			time.Sleep(time.Second) // 避免疯狂重试
			continue
		}

		// 2. 过滤 key
		if "Service" != string(message.Key) {
			continue
		}

		// 3. 反序列化
		var data room_type_data.RoomGuestStay
		if err := json.Unmarshal(message.Value, &data); err != nil {
			zap.L().Error("unmarshal message error", zap.Error(err))
			continue
		}

		// 4. 检查 RPC 客户端
		orderServiceClient := dao.OrderServiceClient
		if orderServiceClient == nil {
			zap.L().Error("OrderService client is nil")
			continue
		}

		// 5. 构造 RPC 请求
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

		rpcCtx, rpcCancel := context.WithTimeout(context.Background(), 2*time.Second)
		_, err = orderServiceClient.OrderInf(rpcCtx, msg)
		rpcCancel()

		if err != nil {
			zap.L().Error(" OrderInf rpc call error ", zap.Error(err))
			continue
		}
	}
}

func SendLog(data []byte) {
	kw.Send(kk.ToSendData{
		Topic: "pms_log",
		Data:  data,
	})
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
