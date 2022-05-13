package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"highroller-go/src/data"
	"highroller-go/src/driver"
)

func CreateUser(user data.User) string {
	var document interface{}
	document = bson.D{
		{"userId", user.UserId},
		{"fullName", user.FullName},
	}
	return driver.CreateOne("user", document)
}

func ReadUser(userId string) data.User {
	return driver.ReadOneUser(userId)
}
