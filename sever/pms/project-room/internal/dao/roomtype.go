package dao

import (
	"context"
	"pms.com/project-room/internal/data/room_type_data"
	"pms.com/project-room/internal/database/gorms"
	"time"
)

type RoomTypeDao struct {
	conn *gorms.GormConn
}

func (r RoomTypeDao) DeleteRoomType(ctx context.Context, roomType *room_type_data.RoomType) error {
	return r.conn.Session(ctx).
		Where("hotel_id = ? AND type_code = ?", roomType.HotelID, roomType.TypeCode).
		Delete(&room_type_data.RoomType{}).Error
}

func (r RoomTypeDao) UpdateRoomType(ctx context.Context, roomType *room_type_data.RoomType) error {
	return r.conn.Session(ctx).
		Model(&room_type_data.RoomType{}).
		Where("hotel_id = ? AND id = ? ", roomType.HotelID, roomType.ID).
		Updates(map[string]any{
			"type_name":     roomType.TypeName,
			"type_code":     roomType.TypeCode,
			"max_occupancy": roomType.MaxOccupancy,
			"base_price":    roomType.BasePrice,
			"quantity":      roomType.Quantity,
			"status":        roomType.Status,
			"updated_at":    time.Now(),
		}).Error
}

func NewRoomTypeDao() *RoomTypeDao {
	return &RoomTypeDao{
		conn: gorms.New(),
	}
}

func (r RoomTypeDao) FindRoomTypes(ctx context.Context, hotelID int64) ([]room_type_data.RoomType, error) {
	var roomTypes []room_type_data.RoomType

	err := r.conn.Session(ctx).
		Where("hotel_id = ? AND is_deleted = ?", hotelID, 0).
		Find(&roomTypes).Error

	return roomTypes, err
}

func (r RoomTypeDao) CreateRoomType(ctx context.Context, roomType *room_type_data.RoomType) error {
	return r.conn.Session(ctx).Create(roomType).Error
}

func (r RoomTypeDao) FindHotelRoom(ctx context.Context, hotelID int64) ([]room_type_data.HotelRoom, error) {
	var hotelRooms []room_type_data.HotelRoom

	err := r.conn.Session(ctx).
		Where("hotel_id = ? ", hotelID).
		Find(&hotelRooms).Error

	return hotelRooms, err
}

func (r RoomTypeDao) CreateHotelRoom(ctx context.Context, room *room_type_data.HotelRoom) error {
	return r.conn.Session(ctx).Create(room).Error
}

func (r RoomTypeDao) UpdateHotelRoom(ctx context.Context, room *room_type_data.HotelRoom) error {
	return r.conn.Session(ctx).
		Model(&room_type_data.HotelRoom{}).
		Where("hotel_id = ? AND id = ? ", room.HotelID, room.ID).
		Updates(map[string]any{
			"room_no":        room.RoomNo,
			"room_type_code": room.RoomTypeCode,
			"room_type_name": room.RoomTypeName,
			"floor_no":       room.FloorNo,
			"phone_ext":      room.PhoneExt,
			"description":    room.Description,
			"updated_at":     time.Now(),
		}).Error
}

func (r RoomTypeDao) DeleteHotelRoom(ctx context.Context, room *room_type_data.HotelRoom) error {
	return r.conn.Session(ctx).
		Where("hotel_id = ? AND room_no = ?", room.HotelID, room.RoomNo).
		Delete(&room_type_data.HotelRoom{}).Error
}

func (r RoomTypeDao) CreateRoomGuestStay(ctx context.Context, room *room_type_data.RoomGuestStay) error {
	return r.conn.Session(ctx).Select("hotel_id", "guest_room_no").Create(room).Error
}

func (r RoomTypeDao) FindRoomGuestStay(ctx context.Context, hotelID int64) ([]room_type_data.RoomGuestStay, error) {
	var roomGuestStays []room_type_data.RoomGuestStay

	err := r.conn.Session(ctx).
		Where("hotel_id = ? ", hotelID).
		Find(&roomGuestStays).Error

	return roomGuestStays, err
}

func (r RoomTypeDao) UpdateRoomGuestStay(ctx context.Context, stay *room_type_data.RoomGuestStay) error {
	return r.conn.Session(ctx).
		Model(&room_type_data.RoomGuestStay{}).
		Where("hotel_id = ? AND guest_room_no = ? ", stay.HotelID, stay.GuestRoomNo).
		Updates(map[string]any{
			"guest_name":     stay.GuestName,
			"guest_id_no":    stay.GuestIDNo,
			"real_price":     stay.RealPrice,
			"mobile":         stay.Mobile,
			"check_in_time":  stay.CheckInTime,
			"check_out_time": stay.CheckOutTime,
			"stay_status":    stay.StayStatus,
			"description":    stay.Description,
		}).Error
}
