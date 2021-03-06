package jwt

import (
	"errors"
	"github.com/spf13/viper"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//过期时间
const TokenExpireDuration = time.Hour * 10000

type TFClaims struct {
	Telephone string `json:"telephone"` // claims 包含一个 telephone
	jwt.StandardClaims
}

//获取token
func GetToken(telephone string) (string, error) {
	//创建声明
	claims := TFClaims{
		telephone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    viper.GetString("server.JwtIssuer"),
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token

	return token.SignedString([]byte(viper.GetString("server.JwtKey")))
}

//解析token
func ParseToken(tokenString string) (*TFClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &TFClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(viper.GetString("server.JwtKey")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TFClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
