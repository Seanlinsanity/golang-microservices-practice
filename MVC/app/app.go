package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

//StartApp is the function of app entry point
func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
