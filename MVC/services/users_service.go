package services

import (
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/domain"
)

//GetUser services
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
