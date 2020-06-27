package app

import (
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/controller/polo"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/controller/repositories"
)

func mapUrls() {
	router.GET("/macro", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
