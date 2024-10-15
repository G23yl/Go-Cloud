package utils

import (
	"disk/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成JWT

type MyClaim struct {
	UserID uint
	jwt.RegisteredClaims
}

var (
	JWTSecret = []byte(config.GetSecret())
	JWTExpire = time.Duration(config.GetExpired()) * time.Hour
)

func GenerateJWT(userID uint) (string, error) {
	myClaim := MyClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWTExpire)),
		},
	}
	// 使用HS256算法生成token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	token, err := claims.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解析JWT
func ParseJWT(tokenString string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token失效")
}