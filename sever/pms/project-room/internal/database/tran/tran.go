package tran

import "pms.com/project-room/internal/database"

type Transaction interface {
	Action(func(conn database.DbConn) error) error
}
