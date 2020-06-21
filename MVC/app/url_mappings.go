package app

import (
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
