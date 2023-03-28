package common

import (
	"myblog/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt加密密钥
var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 用户登录成功后需要为他发放一个token
// 生成token
func ReleaseToken(user model.User) (string, error) {
	// token的有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		//自定义字段
		UserId: user.ID,
		// 标准字段
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发放时间
			IssuedAt: time.Now().Unix(),
		},
	}
	// 使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	// 返回token
	return tokenString, nil
}

// 前端接收到返回的token后会将其保存，当请求需要token验证的接口时再发送给后端
// 后端需要对token进行解析，识别出用户身份
// 解析token的函数
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
