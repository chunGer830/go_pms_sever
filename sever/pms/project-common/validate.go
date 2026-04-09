package common

import "regexp"

var mobileRegexp = regexp.MustCompile(`^1[3-9]\d{9}$`)

func VerifyMobile(mobile string) bool {
	return mobileRegexp.MatchString(mobile)
}
