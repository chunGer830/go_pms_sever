package jwts

import "testing"

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Nzg2OTA0OTIsInRva2VuIjoiMyJ9.Cyf-8o2Y1xGmwMZDQ9JLfQdCVjlZhpndqyXRyhSIqg4"
	ParseToken(tokenString, "pms")
}
