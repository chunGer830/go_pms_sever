package tran

import "pms.com/project-order/internal/database"

type Transaction interface {
	Action(f func(conn database.DbConn) error) error
}
