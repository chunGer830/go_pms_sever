package main

import (
	"github.com/gin-gonic/gin"
	srv "pms.com/project-common"
	_ "pms.com/project-pms/api"
	"pms.com/project-pms/config"
	"pms.com/project-pms/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
