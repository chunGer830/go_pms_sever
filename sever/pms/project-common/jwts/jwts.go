package jwts

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateToken(val string, exp time.Duration, refreshExp time.Duration, secret string, refreshSecret string) *JwtToken {
	aExp := time.Now().Add(exp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
	})
	aToken, _ := accessToken.SignedString([]byte(secret))

	rExp := time.Now().Add(refreshExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   rExp,
	})
	rToken, _ := refreshToken.SignedString([]byte(refreshSecret))

	return &JwtToken{
		AccessToken:  aToken,
		RefreshToken: rToken,
		AccessExp:    aExp,
		RefreshExp:   rExp,
	}
}

func ParseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		val := claims["token"].(string)
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			return "", errors.New("token过期了")
		}
		return val, nil
	} else {
		return "", err
	}
}
