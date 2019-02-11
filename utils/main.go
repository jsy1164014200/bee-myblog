package utils

import (
	"bee-myblog/models"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

var mySigningKey = []byte("sdj&*fds&^DJfd*f92J71!#W#JKHFDhdsf&T&**&*(&")

func EncodeJWT() *models.Token {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    "admin",
		"IssueType": "AccessToken",
		"IssueTime": time.Now().Unix(),
	}).SignedString(mySigningKey)
	if err != nil {
		beego.Error(err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    "admin",
		"IssueType": "RefreshToken",
		"IssueTime": time.Now().Unix(),
	}).SignedString(mySigningKey)
	if err != nil {
		beego.Error(err)
	}
	return &models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func DecodeJWT(tokenStr string) (jwt.MapClaims, error) {
	tokenStruct, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := tokenStruct.Claims.(jwt.MapClaims); ok && tokenStruct.Valid {
		return claim, nil
	} else {
		return nil, errors.New("token is not valid")
	}
}
