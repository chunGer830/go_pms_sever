package model

import (
	"pms.com/project-common/errs"
)

var (
	InLegalMobile = errs.NewError(2001, "手机号不合法")
)
