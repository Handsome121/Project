package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("bala")
	c := MyClaim{
		Username: "xiaomage",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "xiaomage",
		},
	}
	// 加密
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
	}
	// 解密
	token, err := jwt.ParseWithClaims(s, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(token.Claims.(*MyClaim))
}
