package order

import (
	"github.com/gin-gonic/gin"
	"log"
	"pms.com/project-api/api/midd"
	"pms.com/project-api/router"
)

func init() {
	log.Println("init router order")
	ru := &RouterOrder{}
	router.Register(ru)
}

type RouterOrder struct {
}

func (*RouterOrder) Route(r *gin.Engine) {
	//初始化grpc连接
	InitRpcOrderClient()
	h := New()
	group := r.Group("/project/order")
	group.Use(midd.TokenVerify())
	group.POST("/orderInf", h.orderInf)

}
