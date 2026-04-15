package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pms.com/project-api/pkg/model/user"
	common "pms.com/project-common"
	"pms.com/project-common/errs"
	"pms.com/project-grpc/user/login"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	//1.获取参数
	//2.校验参数
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rsp, err := LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})

	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//2. 验证手机合法性
	ctx.JSON(http.StatusOK, rsp.Code)

	/*
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

	*/
}

func (u *HandlerUser) register(c *gin.Context) {
	//1.接收参数
	result := &common.Result{}
	var req user.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//2.校验参数
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	//3.调用grpc 获取响应
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.RegisterMessage{
		HotelName: req.HotelName,
		Username:  req.Username,
		Password:  req.Password,
		Mobile:    req.Mobile,
	}
	_, err = LoginServiceClient.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Sucess("注册成功"))
}

func (u *HandlerUser) login(c *gin.Context) {
	//1.接收参数
	result := &common.Result{}
	var req user.LoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//2. 调用user grpc 登录
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.LoginMessage{
		Username: req.Username,
		Password: req.Password,
	}
	loginRsp, err := LoginServiceClient.Login(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	rsp := &user.LoginRsp{
		Member: user.Member{
			Id:        loginRsp.Member.Id,
			HotelName: loginRsp.Member.HotelName,
		},
		TokenList: user.TokenList{
			AccessToken:    loginRsp.TokenList.AccessToken,
			AccessTokenExp: loginRsp.TokenList.AccessTokenExp,
			TokenType:      loginRsp.TokenList.TokenType,
			RefreshToken:   loginRsp.TokenList.RefreshToken,
		},
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Sucess(rsp))
}

func (u *HandlerUser) changePassword(c *gin.Context) {
	//1.接收参数
	result := &common.Result{}
	var req user.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//2.校验参数
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	//3.调用grpc 获取响应
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.ChangePasswordMessage{
		Username:    req.Username,
		Password:    req.OldPassword,
		NewPassword: req.Password,
	}
	_, err = LoginServiceClient.ChangePassword(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Sucess("修改成功"))
}
