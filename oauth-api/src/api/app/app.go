package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func StartApp() {
	router = gin.Default()

	mapUrls()
	if error := router.Run(":8080"); error != nil {
		panic(error)
	}
}
