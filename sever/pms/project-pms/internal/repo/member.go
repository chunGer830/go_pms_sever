package repo

import "context"

type MemberRepo interface {
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
}
