package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"highroller-go/src/data"
	"highroller-go/src/helper"
	"highroller-go/src/service"
	"io/ioutil"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createUser")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user data.User
	json.Unmarshal(reqBody, &user)
	user.UserId = helper.RandomString(12)
	service.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUser")
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Println("UserId " + userId)
	//service.ReadRoom(roomId)
	json.NewEncoder(w).Encode(service.ReadUser(userId))
}
