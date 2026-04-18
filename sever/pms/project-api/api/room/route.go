package room

import (
	"github.com/gin-gonic/gin"
	"log"
	"pms.com/project-api/router"
)

func init() {
	log.Println("init router user")
	ru := &RouterRoom{}
	router.Register(ru)
}

type RouterRoom struct {
}

func (*RouterRoom) Route(r *gin.Engine) {
	//初始化grpc连接
	InitRpcRoomClient()
	h := New()
	r.POST("/project/room/roomType", h.roomType)
}
