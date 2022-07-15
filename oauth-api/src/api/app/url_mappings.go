package app

import (
	"microservices-ex-app/oauth-api/src/api/controllers/oauth"
	"microservices-ex-app/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
