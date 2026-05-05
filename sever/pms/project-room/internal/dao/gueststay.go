package dao

import (
	"context"
	"pms.com/project-room/internal/data/room_type_data"
	"pms.com/project-room/internal/database/gorms"
)

type RoomGuestStayDao struct {
	conn *gorms.GormConn
}

func NewRoomGuestStayDao() *RoomGuestStayDao {
	return &RoomGuestStayDao{
		conn: gorms.New(),
	}
}

func (r RoomGuestStayDao) CreateRoomGuestStay(ctx context.Context, room *room_type_data.RoomGuestStay) error {
	return r.conn.Session(ctx).Select("hotel_id", "guest_room_no").Create(room).Error
}

func (r RoomGuestStayDao) FindRoomGuestStay(ctx context.Context, hotelID int64) ([]room_type_data.RoomGuestStay, error) {
	var roomGuestStays []room_type_data.RoomGuestStay

	err := r.conn.Session(ctx).
		Where("hotel_id = ? ", hotelID).
		Find(&roomGuestStays).Error

	return roomGuestStays, err
}

func (r RoomGuestStayDao) UpdateRoomGuestStay(ctx context.Context, stay *room_type_data.RoomGuestStay) error {
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
