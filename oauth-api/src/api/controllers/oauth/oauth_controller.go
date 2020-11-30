package oauth

import (
	"github.com/Seanlinsanity/golang-microservices-practice/oauth-api/src/api/domain/oauth"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

}
