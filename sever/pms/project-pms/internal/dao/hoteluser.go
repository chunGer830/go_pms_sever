package dao

import (
	"context"
	"pms.com/project-pms/internal/data/hoteluser"
	"pms.com/project-pms/internal/database/gorms"
)

type HotelUserDao struct {
	conn *gorms.GormConn
}

func (m HotelUserDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	//TODO implement me
	var count int64
	err := m.conn.Session(ctx).Model(&hoteluser.HotelUser{}).Count(&count).Error
	return count > 0, err
}

func NewHotelUserDao() *HotelUserDao {
	return &HotelUserDao{
		conn: gorms.New(),
	}
}
