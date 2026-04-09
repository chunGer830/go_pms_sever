package user

import (
	"errors"
	common "pms.com/project-common"
)

type RegisterReq struct {
	HotelName string `json:"hotel_name" form:"hotel_name"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Mobile    string `json:"mobile" form:"mobile"`
}

func (r RegisterReq) VerifyPassword() bool {
	return r.Password == r.Password2
}

func (r RegisterReq) Verify() error {
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机格式不正确")
	}
	if !r.VerifyPassword() {
		return errors.New("两次密码输入不一致")
	}
	return nil
}
