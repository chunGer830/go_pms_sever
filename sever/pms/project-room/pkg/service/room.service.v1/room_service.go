package room_service_v1

import (
	"context"
	"pms.com/project-grpc/room"
	"pms.com/project-room/internal/dao"
	"pms.com/project-room/internal/repo"
)

type RoomService struct {
	room.UnimplementedRoomServiceServer
	cache repo.Cache
}

func New() *RoomService {
	return &RoomService{
		cache: dao.Rc,
	}
}

func (s *RoomService) RoomType(ctx context.Context, msg *room.RoomTypeMessage) (*room.RoomTypeResponse, error) {
	return &room.RoomTypeResponse{}, nil
}
