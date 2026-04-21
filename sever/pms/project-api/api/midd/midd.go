package midd

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pms.com/project-api/api/user"
	common "pms.com/project-common"
	"pms.com/project-common/errs"
	"pms.com/project-grpc/user/login"
	"time"
)

func TokenVerify() func(c *gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		//1.获取token
		token := c.GetHeader("Authorization")
		//2.token认证
		ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()
		response, err := user.LoginServiceClient.TokenVerify(ctx, &login.LoginMessage{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		//3.处理结果 通过将信息放入gin上下文 未通过返回未登录
		c.Set("hotel_id", response.Member.Id)
		c.Next()
	}
}
