package controllers

import (
	"encoding/json"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/services"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/utils"
	"net/http"
	"strconv"
)

//GetUser request handler
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := (strconv.ParseInt(userIDParam, 10, 64))
	if err != nil {
		//return bad request
		apiErr := &utils.ApplicationError{
			Message:    "user id should be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_requset",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, userErr := services.UsersService.GetUser(userID)
	if userErr != nil {
		jsonValue, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
