package order_service_v1

import (
	"context"
	"go.uber.org/zap"
	"pms.com/project-common/errs"
	"pms.com/project-common/model"
	"pms.com/project-grpc/order/order_inf"
	"pms.com/project-order/internal/dao"
	"pms.com/project-order/internal/data/order_inf_data"
	"pms.com/project-order/internal/database/tran"
	"pms.com/project-order/internal/repo"
	"time"
)

type OrderInfService struct {
	order_inf.UnimplementedOrderServiceServer
	cache        repo.Cache
	orderInfRepo repo.OrderInfRepo
	transaction  tran.Transaction
}

func New() *OrderInfService {
	return &OrderInfService{
		cache:        dao.Rc,
		orderInfRepo: dao.NewOrderInfDao(),
		transaction:  dao.NewTransaction(),
	}
}

func (s *OrderInfService) OrderInf(ctx context.Context, msg *order_inf.OrderInfMessage) (*order_inf.OrderInfResponse, error) {
	checkInTime, err := time.Parse("2006-01-02 15:04:05", msg.CheckInTime)
	if err != nil {
		zap.L().Error("CheckInTime change err ", zap.Error(err))
		return nil, errs.GrpcError(model.TimeChangeError)
	}
	CheckOutTime, err2 := time.Parse("2006-01-02 15:04:05", msg.CheckOutTime)
	if err2 != nil {
		zap.L().Error("CheckOutTime change err ", zap.Error(err))
		return nil, errs.GrpcError(model.TimeChangeError)
	}

	NewOrderInf := &order_inf_data.OrderInfo{
		HotelId:      msg.HotelId,
		GuestRoomNo:  msg.GuestRoomNo,
		GuestName:    msg.GuestName,
		GuestIdNo:    msg.GuestIdNo,
		RealPrice:    uint(msg.RealPrice),
		Mobile:       msg.Mobile,
		CheckInTime:  checkInTime,
		CheckOutTime: CheckOutTime,
	}
	errDb := s.orderInfRepo.SaveOrderInf(ctx, NewOrderInf)
	if errDb != nil {
		zap.L().Error("SaveOrderInf db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	return &order_inf.OrderInfResponse{}, nil
}
