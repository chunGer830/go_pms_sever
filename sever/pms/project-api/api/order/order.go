package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "pms.com/project-common"
	"pms.com/project-common/errs"
	"pms.com/project-grpc/order/order_inf"
	"time"
)

type HandlerOrder struct {
}

func New() *HandlerOrder {
	return &HandlerOrder{}
}

func (r *HandlerOrder) orderInf(c *gin.Context) {
	result := &common.Result{}

	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}

	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &order_inf.OrderInfMessage{
		HotelId: id,
	}

	roomRsp, err := OrderServiceClient.OrderInf(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess(roomRsp))
}
