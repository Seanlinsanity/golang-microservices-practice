package controllers

import (
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/services"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GetUser request handler
func GetUser(c *gin.Context) {
	userID, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))
	if err != nil {
		//return bad request
		apiErr := &utils.ApplicationError{
			Message:    "user id should be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_requset",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
