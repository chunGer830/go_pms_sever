package dao

import (
	"context"
	"gorm.io/gorm"
	"pms.com/project-pms/internal/data/member"
	"pms.com/project-pms/internal/database/gorms"
	"time"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}

func (m *MemberDao) SaveMember(ctx context.Context, men *member.Member) error {
	return m.conn.Session(ctx).Create(men).Error
}

func (m *MemberDao) GetMemberByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("username=?", username).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) FindMember(ctx context.Context, username string, pwd string) (*member.Member, error) {
	var men *member.Member
	err := m.conn.Session(ctx).Where("username=? and password_hash=?", username, pwd).First(&men).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return men, err
}

func (m *MemberDao) GetMemberMessage(ctx context.Context, username string) (*member.Member, error) {
	var memberMessage *member.Member
	err := m.conn.Session(ctx).Where("username = ?", username).First(&memberMessage).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return memberMessage, err
}

func (m *MemberDao) ChangePassword(ctx context.Context, mem *member.Member) error {
	return m.conn.Session(ctx).
		Model(&member.Member{}).
		Where("username = ?", mem.Username).
		Updates(map[string]any{
			"password_hash": mem.PasswordHash,
			"updated_at":    time.Now(),
		}).Error
}
