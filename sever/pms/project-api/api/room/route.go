package room

import (
	"github.com/gin-gonic/gin"
	"log"
	"pms.com/project-api/api/midd"
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
	group := r.Group("/project/room")
	group.Use(midd.TokenVerify())
	group.GET("/roomType", h.roomType)
	group.POST("/saveRoomType", h.saveRoomType)
	group.POST("/updateRoomType", h.updateRoomType)
	group.POST("/deleteRoomType", h.deleteRoomType)

	group.GET("/hotelRoom", h.hotelRoom)
	group.POST("/saveHotelRoom", h.saveHotelRoom)
	group.POST("/updateHotelRoom", h.updateHotelRoom)
	group.POST("/deleteHotelRoom", h.deleteHotelRoom)

	group.GET("/roomGuestStay", h.roomGuestStay)
	group.POST("/updateRoomGuestStay", h.updateRoomGuestStay)
	group.POST("/checkoutRoomGuestStay", h.checkoutRoomGuestStay)
	group.POST("/cleanRoomGuestStay", h.cleanRoomGuestStay)
	group.POST("/disableRoomGuestStay", h.disableRoomGuestStay)
}
