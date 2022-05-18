package token

import (
	"encoding/json"
	"fmt"
	"github.com/aiio/core/config"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// SecretKey JWT SecretKey
var SecretKey = config.V.Jwt.SecretKey

// exp JWT exp
var exp = config.V.Jwt.Exp

// Claims JWT Claims
type Claims struct {
	Exp  int    `json:"exp"`
	Iat  int    `json:"iat"`
	UID  string `json:"uid"`
	Role string `json:"role"`
}

// VerifyToken 验证JWT
func VerifyToken(tokenStr string) (claims Claims, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		return claims, fmt.Errorf("not authorization")
	}
	b, err := json.Marshal(token.Claims)
	if err != nil {
		return claims, err
	}
	jwtClaims := Claims{}
	err = json.Unmarshal(b, &jwtClaims)
	if err != nil {
		return claims, err
	}
	return jwtClaims, nil
}

// GenerateToken 生成JWT
func GenerateToken(uid, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"role": role,
		"exp":  time.Now().Add(time.Minute * time.Duration(exp)).Unix(), // 可以添加过期时间
		"iat":  time.Now().Unix(),
	})
	return token.SignedString([]byte(SecretKey)) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

// RenewToken 续期Token
func RenewToken(tokenStr string) (string, error) {
	claims, err := VerifyToken(tokenStr)
	if err != nil {
		return "", err
	}
	return GenerateToken(claims.UID, claims.Role)
}
