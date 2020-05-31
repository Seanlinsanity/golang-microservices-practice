package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Stephen", LastName: "Curry", Email: "curry@gmail.com"},
	}
)

//GetUser with userID
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user %v was not fount", userID)

}
