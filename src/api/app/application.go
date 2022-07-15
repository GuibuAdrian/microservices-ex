package app

import (
	"github.com/gin-gonic/gin"
	"microservices-ex-app/src/api/log/option_a"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StarApp() {
	option_a.Info("about to start the application", "step:1", "status:pending")
	mapUrls()
	option_a.Info("urls successfully mapped", "step:2", "status:success")

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
