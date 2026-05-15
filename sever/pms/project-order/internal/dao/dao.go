package dao

import (
	"pms.com/project-order/internal/database"
	"pms.com/project-order/internal/database/gorms"
)

type TransactionImpl struct {
	conn database.DbConn
}

func (t TransactionImpl) Action(f func(conn database.DbConn) error) error {
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction() *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.NewTran(),
	}
}
