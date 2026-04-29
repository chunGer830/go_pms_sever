package computer

import (
	"github.com/gin-gonic/gin"
	"log"
	"pms.com/project-api/api/midd"
	"pms.com/project-api/router"
)

func init() {
	log.Println("init router Computer")
	ru := &RouterComputer{}
	router.Register(ru)
}

type RouterComputer struct {
}

func (*RouterComputer) Route(r *gin.Engine) {
	//初始化grpc连接
	InitRpcComputerClient()
	h := New()
	group := r.Group("/project/computer")
	group.Use(midd.TokenVerify())
	group.GET("/computer", h.computer)

}
