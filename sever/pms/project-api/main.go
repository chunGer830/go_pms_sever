package main

import (
	"github.com/gin-gonic/gin"
	_ "pms.com/project-api/api"
	"pms.com/project-api/config"
	"pms.com/project-api/router"
	srv "pms.com/project-common"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
