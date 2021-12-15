package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsonik/go-webapi/handler"
)

// AuthorizedMiddleware verify token
func AuthorizedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := handler.GetUserModel(ctx)
		if err != nil {
			ctx.String(200, err.Error())
			ctx.Abort()
			return
		}
	}
}
