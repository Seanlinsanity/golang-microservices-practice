package oauth

import (
	"strings"

	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (reqeust *AccessTokenRequest) Validate() errors.ApiError {
	reqeust.Username = strings.TrimSpace(reqeust.Username)
	if reqeust.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}

	reqeust.Password = strings.TrimSpace(reqeust.Password)
	if reqeust.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
