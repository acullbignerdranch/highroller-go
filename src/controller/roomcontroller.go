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
	fmt.Println("Endpoint Hit: getRoom")
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	fmt.Println("RoomId " + roomId)
	//service.ReadRoom(roomId)
	json.NewEncoder(w).Encode(service.ReadRoom(roomId))
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getRoom")
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Println("UserId " + userId)
	//service.ReadRoom(roomId)
	json.NewEncoder(w).Encode(service.ReadRooms(userId))
}

//userId: 8F2qNfHK5a84
//roomId: BpLnfg
func joinRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: joinRoom")
	vars := mux.Vars(r)
	userId := vars["userId"]
	//user := service.ReadUser(userId)
	//fmt.Printf("%+v\n", user) // Print with Variable Name
	reqBody, _ := ioutil.ReadAll(r.Body)
	var room data.Room
	//room.Members = append(room.Members, user)
	json.Unmarshal(reqBody, &room)
	room = service.JoinRoom(room, userId)
	json.NewEncoder(w).Encode(room)
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createRoom")
	reqBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	userId := vars["userId"]
	var room data.Room
	json.Unmarshal(reqBody, &room)
	room.HostId = userId
	fmt.Printf("%+v\n", room) // Print with Variable Name
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
