package common

import "regexp"

var mobileRegexp = regexp.MustCompile(`^1[3-9]\d{9}$`)

func VerifyMobile(mobile string) bool {
	return mobileRegexp.MatchString(mobile)
}

func StrVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Int32Val(i *int) int32 {
	if i == nil {
		return 0
	}
	return int32(*i)
}

func StrPtr(s string) *string {
	return &s
}
