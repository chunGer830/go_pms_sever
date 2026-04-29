package dao

import (
	"pms.com/project-computer/internal/database/gorms"
)

type RoomTypeDao struct {
	conn *gorms.GormConn
}

func NewRoomTypeDao() *RoomTypeDao {
	return &RoomTypeDao{
		conn: gorms.New(),
	}
}
