package model

import (
	"pms.com/project-common/errs"
)

var (
	DBError             = errs.NewError(999, "db错误")
	InLegalMobile       = errs.NewError(2001, "手机号不合法")
	UsernameExist       = errs.NewError(102002, "用户名存在")
	UsernameAndPwdError = errs.NewError(102003, "账号或密码错误")
)
