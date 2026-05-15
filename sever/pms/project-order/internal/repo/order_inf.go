package repo

import (
	"context"
	"pms.com/project-order/internal/data/order_inf_data"
)

type OrderInfRepo interface {
	SaveOrderInf(ctx context.Context, orderInf *order_inf_data.OrderInfo) error
}
