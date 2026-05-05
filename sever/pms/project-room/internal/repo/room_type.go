package repo

import (
	"context"
	"pms.com/project-room/internal/data/room_type_data"
)

type RoomTypeRepo interface {
	FindRoomTypes(ctx context.Context, hotelID int64) ([]room_type_data.RoomType, error)
	CreateRoomType(ctx context.Context, roomType *room_type_data.RoomType) error
	UpdateRoomType(ctx context.Context, roomType *room_type_data.RoomType) error
	DeleteRoomType(ctx context.Context, roomType *room_type_data.RoomType) error
	FindHotelRoom(ctx context.Context, hotelID int64) ([]room_type_data.HotelRoom, error)
	CreateHotelRoom(ctx context.Context, room *room_type_data.HotelRoom) error
	UpdateHotelRoom(ctx context.Context, room *room_type_data.HotelRoom) error
	DeleteHotelRoom(ctx context.Context, room *room_type_data.HotelRoom) error
}

type RoomGuestStayRepo interface {
	CreateRoomGuestStay(ctx context.Context, room *room_type_data.RoomGuestStay) error
	FindRoomGuestStay(ctx context.Context, hotelID int64) ([]room_type_data.RoomGuestStay, error)
	UpdateRoomGuestStay(ctx context.Context, stay *room_type_data.RoomGuestStay) error
}
