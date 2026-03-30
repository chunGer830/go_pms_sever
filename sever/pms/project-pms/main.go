package main

import (
	"github.com/gin-gonic/gin"
	"pms.com/project-pms/config"
	"pms.com/project-pms/router"

	srv "pms.com/project-common"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
