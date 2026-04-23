package model

import (
	"pms.com/project-common/errs"
)

var (
	DBError             = errs.NewError(999, "出现错误")
	InLegalMobile       = errs.NewError(2001, "手机号不合法")
	UsernameExist       = errs.NewError(102002, "用户名存在")
	UsernameAndPwdError = errs.NewError(102003, "账号或密码错误")
	NoLogin             = errs.NewError(102004, "未登录")
	NoRoomType          = errs.NewError(102005, "没有房型信息")
	NoHotelRoom         = errs.NewError(102006, "没有房间信息")
	NoRoomGuestStay     = errs.NewError(102006, "没有住客信息")
)
