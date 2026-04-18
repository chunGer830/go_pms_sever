package tran

import "pms.com/project-pms/internal/database"

type Transaction interface {
	Action(func(func(conn database.DbConn)) error) error
}
