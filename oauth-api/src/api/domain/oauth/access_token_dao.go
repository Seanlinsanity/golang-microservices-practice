package oauth

import (
	"fmt"

	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (token *AccessToken) Save() errors.ApiError {
	token.AccessToken = fmt.Sprintf("accessToken_%d", token.UserId)
	tokens[token.AccessToken] = token
	return nil
}

func GetAccessTokenByToken(token string) (*AccessToken, errors.ApiError) {
	aceessToken := tokens[token]
	if aceessToken == nil {
		return nil, errors.NewNotFoundError("invalid token")
	}

	if aceessToken.IsExpired() {
		return nil, errors.NewNotFoundError("token is expired")
	}
	return aceessToken, nil
}
