package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type MyClaims struct {
	UserId   uint
	UserName string
	jwt.RegisteredClaims
}

var (
	// 获取密钥
	secretKey = viper.GetString("jwt.secretKey")
)

// GenerateToken 生成token
func GenerateToken(userId uint, userName string) (string, error) {
	//获取过期时间
	expire := time.Now().Add(time.Duration(viper.GetDuration("jwt.expire") * time.Minute))
	claims := &MyClaims{
		userId,
		userName,
		jwt.RegisteredClaims{
			Subject:   "Token",
			ExpiresAt: jwt.NewNumericDate(expire),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*MyClaims, error) {
	//解析token是否有效
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
