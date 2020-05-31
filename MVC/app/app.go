package app

import (
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/controllers"
	"net/http"
)

//StartApp is the function of app entry point
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
