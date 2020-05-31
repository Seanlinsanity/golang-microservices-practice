package services

import (
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/domain"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/utils"
)

//GetUser services
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
