package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"highroller-go/src/data"
	"highroller-go/src/driver"
)

func CreateRoom(room data.Room) string {
	var document interface{}
	document = bson.D{
		{"roomId", room.RoomId},
		{"hostId", room.HostId},
		{"roomName", room.RoomName},
	}
	return driver.CreateOne("room", document)
}

func ReadRoom(roomId string) data.Room {
	return driver.ReadOneRoom(roomId)
}

func ReadRooms(userId string) []data.Room {
	return driver.ReadManyRooms(userId)
}

func JoinRoom(room data.Room, userId string) data.Room {
	savedRoom := ReadRoom(room.RoomId)
	var savedMembers []string
	if savedRoom.Members != nil {
		savedMembers = append(savedRoom.Members, userId)
	} else {
		savedMembers = append(savedMembers, userId)
	}
	savedRoom.Members = savedMembers
	var update interface{}
	update = bson.D{
		{"$set", bson.D{
			{"members", savedMembers},
		}},
	}
	var filter interface{}
	filter = bson.D{
		{"roomId", room.RoomId},
	}
	result := driver.UpdateOne("room", filter, update)
	if result < 0 {
		fmt.Println("Didn't find any records to update")
	}
	return savedRoom
}
