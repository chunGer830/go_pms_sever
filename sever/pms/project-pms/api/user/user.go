package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	common "pms.com/project-common"
	"pms.com/project-pms/internal/dao"
	"pms.com/project-pms/internal/database/tran"
	"pms.com/project-pms/internal/repo"
	"pms.com/project-pms/pkg/model"
	"pms.com/project-pms/pkg/model/user"
	"time"
)

type HandlerUser struct {
	cache       repo.Cache
	memberRepo  repo.MemberRepo
	transaction tran.Transaction
}

func New() *HandlerUser {
	return &HandlerUser{
		cache:      dao.Rc,
		memberRepo: dao.NewMemberDao(),
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

func (h *HandlerUser) register(ctx *gin.Context) {
	reslut := &common.Result{}
	var req user.RegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, reslut.Fail("error"))
		return
	}
	if err := req.Verify(); err != nil {
		ctx.JSON(http.StatusOK, reslut.Fail("error"))
		return
	}
	//事务操作
	//err := h.tran.Action(func(conn database.DbConn) error {
	//数据库操作
	//数据库操作
	//return nil
	//})
	exist, err := h.memberRepo.GetMemberByEmail(ctx, req.Email)
	if err != nil {
		zap.L().Error("error", zap.Error(err))
		return
	}

}
