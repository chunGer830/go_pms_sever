package main

import (
	"github.com/gin-gonic/gin"
	"pms.com/project-order/config"
	"pms.com/project-order/router"

	srv "pms.com/project-common"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()
	//初始化kafka
	//c := config.InitKafkaWriter()

	stop := func() {
		gc.Stop()
		//c()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
