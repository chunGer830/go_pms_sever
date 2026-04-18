package room

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "pms.com/project-common"
	"pms.com/project-common/errs"
	"pms.com/project-grpc/room"
	"time"
)

type HandlerRoom struct {
}

func New() *HandlerRoom {
	return &HandlerRoom{}
}

func (r *HandlerRoom) roomType(c *gin.Context) {
	//1.接收参数
	result := &common.Result{}

	//2.校验参数
	//3.调用grpc 获取响应
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &room.RoomTypeMessage{}
	_, err := RoomServiceClient.RoomType(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Sucess(""))
}
