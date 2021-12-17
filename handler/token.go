package handler

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xsonik/go-webapi/model"
)

var signingKey = []byte("05905F9A-9CFE-40A7-8907-27C52C0E98D1")

func GetToken(ctx *gin.Context) {
	claims := &model.MyClaims{}
	err := ctx.BindJSON(claims)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	expiresAt := time.Now().Add(time.Minute * 30)
	claims.Subject = "webapi token"
	claims.Issuer = "webapi"
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = expiresAt.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		ctx.String(500, err.Error())
		return
	}
	ctx.JSON(200, gin.H{
		"user_code": claims.UserCode,
		"token":     tokenString,
		"expired":   expiresAt,
	})
}

func GetUserModel(ctx *gin.Context) (model.UserInfo, error) {
	userinfo := model.UserInfo{}
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		return userinfo, errors.New("token is empty")
	}
	if len(tokenString) > 6 && strings.ToLower(tokenString[:6]) == "bearer" {
		tokenString = tokenString[7:]
	}
	claims := &model.MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return userinfo, err
	}
	if !token.Valid {
		return userinfo, errors.New("token is invalid")
	}
	userinfo = model.UserInfo{
		UserCode:   claims.UserCode,
		UserName:   claims.UserName,
		Filiale:    claims.Filiale,
		ModuleFlag: claims.ModuleFlag,
		OperHost:   claims.OperHost,
		OperIP:     claims.OperIP,
	}
	return userinfo, nil
}
