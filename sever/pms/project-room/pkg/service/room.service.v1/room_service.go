package room_service_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	common "pms.com/project-common"
	"pms.com/project-common/errs"
	"pms.com/project-common/model"
	"pms.com/project-grpc/room/room_type"
	"pms.com/project-room/internal/dao"
	"pms.com/project-room/internal/data/room_type_data"
	"pms.com/project-room/internal/database/tran"
	"pms.com/project-room/internal/repo"
	"time"
)

type RoomService struct {
	room_type.UnimplementedRoomServiceServer
	cache             repo.Cache
	roomTypeRepo      repo.RoomTypeRepo
	roomGuestStayRepo repo.RoomGuestStayRepo
	transaction       tran.Transaction
}

func New() *RoomService {
	return &RoomService{
		cache:             dao.Rc,
		roomTypeRepo:      dao.NewRoomTypeDao(),
		roomGuestStayRepo: dao.NewRoomTypeDao(),
	}
}

func (s *RoomService) RoomType(ctx context.Context, msg *room_type.RoomTypeMessage) (*room_type.RoomTypeResponse, error) {
	//读缓存
	key := fmt.Sprintf("%s%d", model.RoomTypeRedis, msg.HotelId)
	val, err := s.cache.Get(ctx, key)
	if err == nil && val != "" {
		var rsp room_type.RoomTypeResponse
		if err := json.Unmarshal([]byte(val), &rsp); err == nil {
			return &rsp, nil
		}
	}
	//读数据库
	roomTypes, err := s.roomTypeRepo.FindRoomTypes(ctx, msg.HotelId)
	if err != nil {
		zap.L().Error("RoomType db FindRoomTypes err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	if len(roomTypes) == 0 {
		return nil, errs.GrpcError(model.NoRoomType)
	}

	list := make([]*room_type.RoomTypeItem, 0, len(roomTypes))
	for _, rt := range roomTypes {

		list = append(list, &room_type.RoomTypeItem{
			Id:           rt.ID,
			TypeName:     rt.TypeName,
			TypeCode:     rt.TypeCode,
			MaxOccupancy: int32(rt.MaxOccupancy),
			BasePrice:    rt.BasePrice,
			Quantity:     int32(rt.Quantity),
			Status:       int32(rt.Status),
		})
	}

	rsp := &room_type.RoomTypeResponse{
		List: list,
	}

	//存缓存
	data, err := json.Marshal(rsp)
	if err == nil {
		cacheErr := s.cache.Put(ctx, key, string(data), 30*time.Minute)
		if cacheErr != nil {
			zap.L().Error("RoomType cache put err", zap.Error(cacheErr))
		}
	} else {
		zap.L().Error("RoomType json marshal err", zap.Error(err))
	}

	return rsp, nil
}

func (s *RoomService) SaveRoomType(ctx context.Context, msg *room_type.SaveRoomTypeMessage) (*room_type.SaveRoomTypeResponse, error) {
	NewRoomType := &room_type_data.RoomType{
		HotelID:      msg.HotelId,
		TypeCode:     msg.TypeCode,
		TypeName:     msg.TypeName,
		MaxOccupancy: int(msg.MaxOccupancy),
		BasePrice:    msg.BasePrice,
		Quantity:     int(msg.Quantity),
		Status:       int(msg.Status),
	}
	err := s.roomTypeRepo.CreateRoomType(ctx, NewRoomType)
	if err != nil {
		zap.L().Error("CreateRoomType db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存
	key := fmt.Sprintf("%s%d", model.RoomTypeRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key)

	return &room_type.SaveRoomTypeResponse{}, nil
}

func (s *RoomService) UpdateRoomType(ctx context.Context, msg *room_type.UpdateRoomTypeMessage) (*room_type.UpdateRoomTypeResponse, error) {
	NewRoomType := &room_type_data.RoomType{
		ID:           msg.Id,
		HotelID:      msg.HotelId,
		TypeCode:     msg.TypeCode,
		TypeName:     msg.TypeName,
		MaxOccupancy: int(msg.MaxOccupancy),
		BasePrice:    msg.BasePrice,
		Quantity:     int(msg.Quantity),
		Status:       int(msg.Status),
	}
	err := s.roomTypeRepo.UpdateRoomType(ctx, NewRoomType)
	if err != nil {
		zap.L().Error("UpdateRoomType db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存
	key := fmt.Sprintf("%s%d", model.RoomTypeRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key)

	return &room_type.UpdateRoomTypeResponse{}, nil
}

func (s *RoomService) DeleteRoomType(ctx context.Context, msg *room_type.DeleteRoomTypeMessage) (*room_type.DeleteRoomTypeResponse, error) {
	NewRoomType := &room_type_data.RoomType{
		HotelID:  msg.HotelId,
		TypeCode: msg.TypeCode,
	}
	err := s.roomTypeRepo.DeleteRoomType(ctx, NewRoomType)
	if err != nil {
		zap.L().Error("DeleteRoomType db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存
	key := fmt.Sprintf("%s%d", model.RoomTypeRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key)

	key2 := fmt.Sprintf("%s%d", model.HotelRoomRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key2)

	key3 := fmt.Sprintf("%s%d", model.RoomGuestStayRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key3)

	return &room_type.DeleteRoomTypeResponse{}, nil
}

func (s *RoomService) HotelRoom(ctx context.Context, msg *room_type.HotelRoomMessage) (*room_type.HotelRoomResponse, error) {
	//查缓存
	key := fmt.Sprintf("%s%d", model.HotelRoomRedis, msg.HotelId)
	val, err := s.cache.Get(ctx, key)
	if err == nil && val != "" {
		var rsp room_type.HotelRoomResponse
		if err := json.Unmarshal([]byte(val), &rsp); err == nil {
			return &rsp, nil
		}
	}
	//查数据库
	HotelRooms, err := s.roomTypeRepo.FindHotelRoom(ctx, msg.HotelId)
	if err != nil {
		zap.L().Error("HotelRoom db FindHotelRoom err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	if len(HotelRooms) == 0 {
		return nil, errs.GrpcError(model.NoHotelRoom)
	}

	list := make([]*room_type.HotelRoomItem, 0, len(HotelRooms))
	for _, rt := range HotelRooms {

		list = append(list, &room_type.HotelRoomItem{
			Id:           rt.ID,
			RoomNo:       rt.RoomNo,
			RoomTypeName: rt.RoomTypeName,
			FloorNo:      common.StrVal(rt.FloorNo),
			Building:     common.StrVal(rt.Building),
			PhoneExt:     common.StrVal(rt.PhoneExt),
			Description:  common.StrVal(rt.Description),
		})
	}

	rsp := &room_type.HotelRoomResponse{
		List: list,
	}

	//存缓存
	data, err := json.Marshal(rsp)
	if err == nil {
		cacheErr := s.cache.Put(ctx, key, string(data), 30*time.Minute)
		if cacheErr != nil {
			zap.L().Error("HotelRoom cache put err", zap.Error(cacheErr))
		}
	} else {
		zap.L().Error("HotelRoom json marshal err", zap.Error(err))
	}

	return rsp, nil
}

func (s *RoomService) SaveHotelRoom(ctx context.Context, msg *room_type.SaveHotelRoomMessage) (*room_type.SaveHotelRoomResponse, error) {
	NewHotelRoom := &room_type_data.HotelRoom{
		HotelID:      msg.HotelId,
		RoomNo:       msg.RoomNo,
		RoomTypeName: msg.RoomTypeName,
		RoomTypeCode: msg.RoomTypeCode,
		FloorNo:      common.StrPtr(msg.FloorNo),
		PhoneExt:     common.StrPtr(msg.PhoneExt),
		Description:  common.StrPtr(msg.Description),
	}

	NewRoomGuestStay := &room_type_data.RoomGuestStay{
		HotelID:     msg.HotelId,
		GuestRoomNo: msg.RoomNo,
	}

	err := s.roomTypeRepo.CreateHotelRoom(ctx, NewHotelRoom)
	_ = s.roomTypeRepo.CreateRoomGuestStay(ctx, NewRoomGuestStay)
	if err != nil {
		zap.L().Error("CreateHotelRoom db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存
	key2 := fmt.Sprintf("%s%d", model.HotelRoomRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key2)

	key3 := fmt.Sprintf("%s%d", model.RoomGuestStayRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key3)

	return &room_type.SaveHotelRoomResponse{}, nil
}

func (s *RoomService) UpdateHotelRoom(ctx context.Context, msg *room_type.UpdateHotelRoomMessage) (*room_type.UpdateHotelRoomResponse, error) {
	NewHotelRoom := &room_type_data.HotelRoom{
		ID:           msg.Id,
		HotelID:      msg.HotelId,
		RoomNo:       msg.RoomNo,
		RoomTypeName: msg.RoomTypeName,
		RoomTypeCode: msg.RoomTypeCode,
		FloorNo:      common.StrPtr(msg.FloorNo),
		PhoneExt:     common.StrPtr(msg.PhoneExt),
		Description:  common.StrPtr(msg.Description),
	}
	err := s.roomTypeRepo.UpdateHotelRoom(ctx, NewHotelRoom)
	if err != nil {
		zap.L().Error("UpdateHotelRoom db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存
	key2 := fmt.Sprintf("%s%d", model.HotelRoomRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key2)

	key3 := fmt.Sprintf("%s%d", model.RoomGuestStayRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key3)

	return &room_type.UpdateHotelRoomResponse{}, nil
}

func (s *RoomService) DeleteHotelRoom(ctx context.Context, msg *room_type.DeleteHotelRoomMessage) (*room_type.DeleteHotelRoomResponse, error) {
	NewHotelRoom := &room_type_data.HotelRoom{
		HotelID: msg.HotelId,
		RoomNo:  msg.RoomNo,
	}
	err := s.roomTypeRepo.DeleteHotelRoom(ctx, NewHotelRoom)
	if err != nil {
		zap.L().Error("DeleteHotelRoom db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存
	key2 := fmt.Sprintf("%s%d", model.HotelRoomRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key2)
	key3 := fmt.Sprintf("%s%d", model.RoomGuestStayRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key3)

	return &room_type.DeleteHotelRoomResponse{}, nil
}

func (s *RoomService) RoomGuestStay(ctx context.Context, msg *room_type.RoomGuestStayMessage) (*room_type.RoomGuestStayResponse, error) {
	//读缓存
	key := fmt.Sprintf("%s%d", model.RoomGuestStayRedis, msg.HotelId)
	val, err := s.cache.Get(ctx, key)
	if err == nil && val != "" {
		var rsp room_type.RoomGuestStayResponse
		if err := json.Unmarshal([]byte(val), &rsp); err == nil {
			return &rsp, nil
		}
	}
	//读数据库
	RoomGuestStays, err := s.roomGuestStayRepo.FindRoomGuestStay(ctx, msg.HotelId)
	if err != nil {
		zap.L().Error("RoomGuestStay db FindRoomGuestStay err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	if len(RoomGuestStays) == 0 {
		return nil, errs.GrpcError(model.NoRoomGuestStay)
	}

	list := make([]*room_type.RoomGuestStayItem, 0, len(RoomGuestStays))
	for _, rt := range RoomGuestStays {

		list = append(list, &room_type.RoomGuestStayItem{
			Id:           rt.ID,
			HotelId:      rt.HotelID,
			GuestRoomNo:  rt.GuestRoomNo,
			GuestName:    rt.GuestName,
			GuestIdNo:    rt.GuestIDNo,
			RealPrice:    rt.RealPrice,
			Mobile:       rt.Mobile,
			CheckInTime:  rt.CheckInTime,
			CheckOutTime: rt.CheckOutTime,
			StayStatus:   int32(rt.StayStatus),
			Description:  rt.Description,
		})
	}

	rsp := &room_type.RoomGuestStayResponse{
		List: list,
	}

	//存缓存
	data, err := json.Marshal(rsp)
	if err == nil {
		cacheErr := s.cache.Put(ctx, key, string(data), 30*time.Minute)
		if cacheErr != nil {
			zap.L().Error("RoomGuestStay cache put err", zap.Error(cacheErr))
		}
	} else {
		zap.L().Error("RoomGuestStay json marshal err", zap.Error(err))
	}

	return rsp, nil
}

func (s *RoomService) UpdateRoomGuestStay(ctx context.Context, msg *room_type.UpdateRoomGuestStayMessage) (*room_type.UpdateRoomGuestStayResponse, error) {
	NewRoomGuestStay := &room_type_data.RoomGuestStay{
		ID:           msg.Id,
		HotelID:      msg.HotelId,
		GuestRoomNo:  msg.GuestRoomNo,
		GuestName:    msg.GuestName,
		GuestIDNo:    msg.GuestIdNo,
		RealPrice:    msg.RealPrice,
		Mobile:       msg.Mobile,
		CheckInTime:  msg.CheckInTime,
		CheckOutTime: msg.CheckOutTime,
		StayStatus:   int8(msg.StayStatus),
		Description:  msg.Description,
	}
	err := s.roomGuestStayRepo.UpdateRoomGuestStay(ctx, NewRoomGuestStay)
	if err != nil {
		zap.L().Error("UpdateRoomGuestStay db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}

	//删缓存

	key3 := fmt.Sprintf("%s%d", model.RoomGuestStayRedis, msg.HotelId)
	_ = s.cache.Delete(ctx, key3)

	return &room_type.UpdateRoomGuestStayResponse{}, nil
}
