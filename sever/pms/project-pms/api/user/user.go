package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"pms.com/project-pms/pkg/dao"
	"pms.com/project-pms/pkg/model"
	"pms.com/project-pms/pkg/repo"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	//1.获取参数
	//2.校验参数
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := h.cache.Put(c, "1", "2", 15*time.Minute)
	if err != nil {
		zap.L().Error("redis put error", zap.String("error", err.Error()))
	}
	zap.L().Info("存入redis成功")
	//zap.L().Error("redis error")
	val, err := h.cache.Get(c, "1")
	fmt.Println(val)
	ctx.JSON(http.StatusOK, model.InLegalMobile)
}
