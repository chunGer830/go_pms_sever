package login_service_v1

import (
	"context"
	"fmt"
	"log"
	common "pms.com/project-common"
	"pms.com/project-pms/internal/dao"
	"pms.com/project-pms/internal/repo"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (h *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	rsq := &common.Result{}
	fmt.Println(rsq)
	//1. 获取参数
	mobile := msg.Mobile
	//2. 验证手机合法性
	//if {
	//	return nil,errs.GrpcError(model.InLegalMobile)
	//}

	//3.生成验证码
	code := "123456"
	//4. 发送验证码
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("调用短信平台发送短信")
		//发送成功 存入redis
		fmt.Println(mobile, code)
	}()
	return &CaptchaResponse{Code: code}, nil
}
