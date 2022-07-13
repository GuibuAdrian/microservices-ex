package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	//router = gin.New()
}

func StartApp() {
	mapUrls()

	if err := router.Run("localhost:8084"); err != nil {
		panic(err)
	}
}
