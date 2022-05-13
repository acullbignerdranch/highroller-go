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
)

func HandleRouters() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/{userId}/room", getRooms).Methods("GET")
	myRouter.HandleFunc("/user/{userId}/room", createRoom).Methods("POST")
	myRouter.HandleFunc("/user/{userId}/room", joinRoom).Methods("PUT")
	myRouter.HandleFunc("/user/{userId}", getUser).Methods("GET")
	myRouter.HandleFunc("/user/{userId}/room", deleteRoom).Methods("DELETE")
	myRouter.HandleFunc("/user", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	json.NewEncoder(w).Encode(service.ReadRoom(roomId))
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	json.NewEncoder(w).Encode(service.ReadRooms(userId))
}

func joinRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var room data.Room
	json.Unmarshal(reqBody, &room)
	room = service.JoinRoom(room, userId)
	json.NewEncoder(w).Encode(room)
}

func createRoom(w http.ResponseWriter, r *http.Request) {
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

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	roomId := vars["id"]

	//TODO: delete the rooms from the database
	fmt.Println(roomId)
}
