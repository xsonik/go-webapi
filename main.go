package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xsonik/go-webapi/router"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	appEngine := gin.Default()
	router.RegisterRoutes(appEngine)

	log.Println("Starting Webapi server...")
	err := appEngine.Run(":8056")
	if err != nil {
		log.Print(err.Error())
	}
}
