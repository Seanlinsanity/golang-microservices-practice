package controllers

import (
	"encoding/json"
	"github.com/Seanlinsanity/golang-microservices-practice/MVC/services"
	"net/http"
	"strconv"
)

//GetUser request handler
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := (strconv.ParseInt(userIDParam, 10, 64))
	if err != nil {
		//return bad request
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user id should be a number"))
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		//Handle the err
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
