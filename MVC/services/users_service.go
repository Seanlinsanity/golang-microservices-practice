package services

import (
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/domain"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/utils"
)

type usersService struct {
}

//UserService instance
var (
	UsersService usersService
)

//GetUser services
func (service *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID)
}
