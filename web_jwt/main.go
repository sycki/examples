package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var (
	signingKey = []byte("abc")
)

func main() {
	tokenStr, err := GenToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenStr)
	user, err := AuthenticationToken(tokenStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

// GenToken 生成Token
func GenToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	t := time.Now()
	// 经测试，Token中可以存入任意多的数据，生成的Token串也会相应的变长
	token.Claims = &jwt.StandardClaims{
		Subject:   "jack",
		ExpiresAt: t.Add(time.Minute * 30).Unix(),
		IssuedAt:  t.Unix(),
	}
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		err = errors.Wrap(err, "token.SignedString")
		log.Errorln(err)
		return "", err
	}
	return tokenString, nil
}

// AuthenticationToken 验证Token信息是否有效
func AuthenticationToken(tokenStr string) (*jwt.StandardClaims, error) {
	user := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, user, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		err := errors.Wrap(err, "jwt.Parse")
		return nil, err
	}

	if !token.Valid {
		err := errors.New("Invalid token")
		return nil, err
	}

	return user, nil
}
