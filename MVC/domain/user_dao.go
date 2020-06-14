package domain

import (
	"fmt"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Stephen", LastName: "Curry", Email: "curry@gmail.com"},
	}

	//UserDao user service interface
	UserDao usersServiceInterface
)

func init() {
	UserDao = &userDao{}
}

type usersServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

//GetUser with userID
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing database")
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not fount", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
