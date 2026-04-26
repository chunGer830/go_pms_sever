package login_service_v1

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	common "pms.com/project-common"
	"pms.com/project-common/encrypts"
	"pms.com/project-common/errs"
	"pms.com/project-common/jwts"
	"pms.com/project-common/model"
	"pms.com/project-grpc/user/login"
	"pms.com/project-pms/config"
	"pms.com/project-pms/internal/dao"
	"pms.com/project-pms/internal/data/member"
	"pms.com/project-pms/internal/repo"
	"strconv"
	"time"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache      repo.Cache
	memberRepo repo.MemberRepo
}

func New() *LoginService {
	return &LoginService{
		cache:      dao.Rc,
		memberRepo: dao.NewMemberDao(),
	}
}

func (h *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	//校验 （账号是否被注册）
	c := context.Background()
	exist, err := h.memberRepo.GetMemberByUsername(c, msg.Username)
	if err != nil {
		zap.L().Error("Register db get err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.UsernameExist)
	}
	//存入数据
	pwd := encrypts.Md5(msg.Password)
	mem := &member.Member{
		HotelName:    msg.HotelName,
		Username:     msg.Username,
		PasswordHash: pwd,
		Mobile:       &msg.Mobile,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Status:       1,
	}
	err = h.memberRepo.SaveMember(c, mem)
	if err != nil {
		zap.L().Error("Register db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	//返回
	return &login.RegisterResponse{}, nil
}

func (h *LoginService) Login(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	c := context.Background()
	//1.数据库查询账号密码是否正确
	pwd := encrypts.Md5(msg.Password)
	mem, err := h.memberRepo.FindMember(c, msg.Username, pwd)
	if err != nil {
		zap.L().Error("Login db FindMember err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if mem == nil {
		return nil, errs.GrpcError(model.UsernameAndPwdError)
	}
	//查询hotel_name
	idAndhotel_name, err := h.memberRepo.GetMemberMessage(c, msg.Username)
	if err != nil {
		zap.L().Error("Login db GetMemberMessage err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	memberMessage := &login.MemberMessage{
		Id:        idAndhotel_name.ID,
		HotelName: idAndhotel_name.HotelName,
	}
	//2.jwt生成token
	memIdStr := strconv.FormatInt(idAndhotel_name.ID, 10)
	exp := time.Duration(config.C.JwtConfig.AccessExp) * 24 * time.Hour
	refreshExp := time.Duration(config.C.JwtConfig.RefreshExp) * 24 * time.Hour
	token := jwts.CreateToken(memIdStr, exp, refreshExp, config.C.JwtConfig.AccessSecret, config.C.JwtConfig.RefreshSecret)
	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bearer",
		AccessTokenExp: token.AccessExp,
	}
	return &login.LoginResponse{
		Member:    memberMessage,
		TokenList: tokenList,
	}, nil
}

func (h *LoginService) ChangePassword(ctx context.Context, msg *login.ChangePasswordMessage) (*login.ChangePasswordResponse, error) {
	c := context.Background()
	//1.数据库查询账号密码是否正确
	pwd := encrypts.Md5(msg.Password)
	mem, err := h.memberRepo.FindMember(c, msg.Username, pwd)
	if err != nil {
		zap.L().Error("ChangePassword db FindMember err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if mem == nil {
		return nil, errs.GrpcError(model.UsernameAndPwdError)
	}

	//存入数据
	newPwd := encrypts.Md5(msg.NewPassword)
	mem = &member.Member{
		Username:     msg.Username,
		PasswordHash: newPwd,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Status:       1,
	}
	err = h.memberRepo.ChangePassword(c, mem)
	if err != nil {
		zap.L().Error("ChangePassword db save err ", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	//返回
	return &login.ChangePasswordResponse{}, nil
}

func (h *LoginService) TokenVerify(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	token := msg.Token
	parseToken, err := jwts.ParseToken(token, config.C.JwtConfig.AccessSecret)
	if err != nil {
		zap.L().Error("Login TokenVerify error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	id, _ := strconv.ParseInt(parseToken, 10, 64)
	memberMessage := &login.MemberMessage{
		Id: id,
	}
	return &login.LoginResponse{
		Member: memberMessage,
	}, nil
}
