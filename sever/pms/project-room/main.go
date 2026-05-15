package main

import (
	"github.com/gin-gonic/gin"
	"pms.com/project-room/config"
	"pms.com/project-room/internal/consumer"
	"pms.com/project-room/internal/dao"
	"pms.com/project-room/router"

	srv "pms.com/project-common"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()

	//初始化kafka生产者
	c := consumer.InitKafkaWriter()
	//初始化kafka消费者
	reader := consumer.NewOrderServiceReader()
	go reader.OrderService()

	//
	dao.InitRpcOrderClient()

	stop := func() {
		gc.Stop()
		c()
		reader.R.Close()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
