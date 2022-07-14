package app

import (
	"microservices-ex-app/src/api/controllers/polo"
	"microservices-ex-app/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
