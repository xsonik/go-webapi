package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsonik/go-webapi/handler"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", handler.Default)
	router.GET("/api", handler.Default)

	auth := router.Group("/api/auth")
	{
		auth.POST("", handler.Default)
		auth.POST("token", handler.GetToken)
		auth.Use(AuthorizedMiddleware())
		auth.POST("userinfo", handler.GetUserInfo)
	}
}
