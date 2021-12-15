package handler

import (
	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	root := ctx.FullPath()
	result := "%s it's work!"
	ctx.Header("Content-Type", "text/html")
	ctx.String(200, result, root[1:])
}

func GetUserInfo(ctx *gin.Context) {
	usr, err := GetUserModel(ctx)
	if err != nil {
		ctx.String(200, err.Error())
		return
	}
	ctx.JSON(200, usr)
}
