package dao

import (
	"context"
	"pms.com/project-order/internal/data/order_inf_data"
	"pms.com/project-order/internal/database/gorms"
)

type OrderInfDao struct {
	conn *gorms.GormConn
}

func NewOrderInfDao() *OrderInfDao {
	return &OrderInfDao{
		conn: gorms.New(),
	}
}

func (o OrderInfDao) SaveOrderInf(ctx context.Context, orderInf *order_inf_data.OrderInfo) error {
	return o.conn.Session(ctx).Create(orderInf).Error
}
