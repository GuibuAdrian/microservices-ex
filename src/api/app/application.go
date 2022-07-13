package app

import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StarApp() {
	mapUrls()
	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
