package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	// 加密的key
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	// Create the Claims
	claims := MyCustomClaims{
		"超越",
		jwt.StandardClaims{
			// 过期时间(5秒后过期)
			ExpiresAt: time.Now().Unix() + 5,
			// 签发人
			Issuer: "test",
			// 生效时间，可配置可不配置
			// NotBefore: time.Now().Unix() - 5,
		},
	}
	// 参数：加密方法；实现 claims 接口的结构体
	// claims 可以替换成 jwt.MapClaims{}使用，最终代码为 -> token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": "chaoyue"...})

	// 获取完整的签名令牌
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("加密失败， err:", err)
	}
	fmt.Printf("%v | %v\n", ss, err)

	// time.Sleep(6 * time.Second)
	// 解密
	token, err = jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	fmt.Println(token.Claims.(*MyCustomClaims).Foo, err)
}
