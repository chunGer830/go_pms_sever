package tran

import "pms.com/project-computer/internal/database"

type Transaction interface {
	Action(func(func(conn database.DbConn)) error) error
}
