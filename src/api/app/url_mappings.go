package app

import (
	"microservices-ex-app/src/api/controllers/polo"
	"microservices-ex-app/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
