package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"pms.com/project-api/router"
)

func init() {
	log.Println("init router user")
	ru := &RouterUser{}
	router.Register(ru)
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	//初始化grpc连接
	InitRpcUserClient()
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	//r.POST("/project/login/register", h.register)
}
