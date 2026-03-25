package dao

import (
	"context"
	"pms.com/project-pms/internal/data/member"
	"pms.com/project-pms/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	//TODO implement me
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Count(&count).Error
	return count > 0, err
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}
