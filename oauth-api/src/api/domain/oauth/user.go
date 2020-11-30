package oauth

import (
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

const (
	queryGetUser = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"curry": {Id: 30, Username: "Stphen Curry"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user exist")
	}

	return user, nil
}
