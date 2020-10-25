package util

import (
	"JalanJiang/douban-tool-server/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

// Claims JWT 数据封装
type Claims struct {
	UserID uint   `json:"user_id"`
	OpenID string `json:"open_id"`
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(userID uint, openID string) (string, error) {
	nowTime := time.Now()
	// 3 小时过期时间
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		userID,
		openID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "dou-push-server", // 签发人
		},
	}

	// 加密生成 Token claims
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串 Token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析 Token，返回 Claims 结构体
func ParseToken(token string) (*Claims, error) {
	// 解析鉴权声明，内部解码和校验
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 校验成功
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
