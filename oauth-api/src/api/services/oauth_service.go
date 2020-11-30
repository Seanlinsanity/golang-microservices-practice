package services

import (
	"github.com/Seanlinsanity/golang-microservices-practice/oauth-api/src/api/domain/oauth"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (oauth.AccessTokenRequest, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}
