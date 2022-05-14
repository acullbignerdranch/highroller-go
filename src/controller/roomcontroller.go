package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"highroller-go/src/data"
	"highroller-go/src/helper"
	"highroller-go/src/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HandleRouters() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/{userId}/room", getRooms).Methods("GET")
	myRouter.HandleFunc("/user/{userId}/room", createRoom).Methods("POST")
	myRouter.HandleFunc("/user/{userId}/room", joinRoom).Methods("PUT")
	myRouter.HandleFunc("/user/{userId}", getUser).Methods("GET")
	myRouter.HandleFunc("/user/{userId}/room", deleteRoom).Methods("DELETE")
	myRouter.HandleFunc("/user/{userId}/roll", saveRoll).Methods("PUT")
	myRouter.HandleFunc("/user", createUser).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" // Default port if not specified
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	json.NewEncoder(w).Encode(service.ReadRoom(roomId))
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	json.NewEncoder(w).Encode(service.ReadRooms(userId))
}

func joinRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var room data.Room
	json.Unmarshal(reqBody, &room)
	room = service.JoinRoom(room, userId)
	json.NewEncoder(w).Encode(room)
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	userId := vars["userId"]
	var room data.Room
	json.Unmarshal(reqBody, &room)
	room.HostId = userId
	room.RoomId = helper.RandomString(6)
	room.DocId = service.CreateRoom(room)
	json.NewEncoder(w).Encode(room)
}

func saveRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	userId := vars["userId"]
	var roll data.Roll
	json.Unmarshal(reqBody, &roll)
	roll.UserId = userId
	room := service.SaveRoll(roll)
	json.NewEncoder(w).Encode(room)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	roomId := vars["id"]

	//TODO: delete the rooms from the database
	fmt.Println(roomId)
}

/**
userRolls[{
userid
roomid
rolls[{[dice]}]
}]
*/
