package main

import (
	"github.com/gin-gonic/gin"
	"pms.com/project-computer/config"
	"pms.com/project-computer/router"

	srv "pms.com/project-common"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()

	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
