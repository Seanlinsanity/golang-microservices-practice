package domain

import (
	"fmt"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Stephen", LastName: "Curry", Email: "curry@gmail.com"},
	}
)

//GetUser with userID
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not fount", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
