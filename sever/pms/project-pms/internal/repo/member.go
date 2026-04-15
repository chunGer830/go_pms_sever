package repo

import (
	"context"
	"pms.com/project-pms/internal/data/member"
)

type MemberRepo interface {
	GetMemberByUsername(ctx context.Context, username string) (bool, error)
	SaveMember(ctx context.Context, men *member.Member) error
	FindMember(ctx context.Context, username string, pwd string) (*member.Member, error)
	GetMemberMessage(ctx context.Context, username string) (*member.Member, error)
	ChangePassword(ctx context.Context, mem *member.Member) error
}
