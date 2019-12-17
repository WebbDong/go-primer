package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	key := "c3dlbGwtYXBwLTIwMTgtMTAtMjktMDAlM0EyMyUzQTIyJTIwZnVja2luZyUyMHNoaXQ="
	claims := jwt.MapClaims{}
	claims["email"] = "32432432@qq.com"
	claims["pwd"] = "1111111111"
	//claims["expire"] =
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(signedString)

	token, err = jwt.Parse(signedString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token.Valid)
	fmt.Println(token.Claims)
}
