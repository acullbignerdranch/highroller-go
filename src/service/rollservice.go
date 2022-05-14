package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"highroller-go/src/data"
	"highroller-go/src/driver"
	"strconv"
)

func SaveRoll(roll data.Roll) *data.Room {
	var room = driver.ReadOneRoom(roll.RoomId)
	if room.Rolls == nil {
		room.Rolls = make(map[string]int)
	}
	value, ok := room.Rolls[roll.UserId]
	if !ok {
		room.Rolls[roll.UserId] = roll.Roll
		var update interface{}
		update = bson.D{
			{"$set", bson.D{
				{"rolls", room.Rolls},
			}},
		}
		var filter interface{}
		filter = bson.D{
			{"roomId", room.RoomId},
		}
		driver.UpdateOne("room", filter, update)
	} else {
		fmt.Println("Roll already found:" + strconv.Itoa(value))
	}
	return &room
}
