package app

import (
	"github.com/Seanlinsanity/golang-microservices-practice/oauth-api/src/api/controllers/oauth"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/controller/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
