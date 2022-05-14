package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"highroller-go/src/data"
	"log"
	"os"
	"time"
)

var dataBaseUri = os.Getenv("MONGODB_URI")
var databaseName = "highroller"

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		5*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func CreateOne(col string, doc interface{}) string {
	client, ctx, cancel, err := Connect(dataBaseUri)
	if err != nil {
		panic(err)
	}
	defer Close(client, ctx, cancel)

	collection := client.Database(databaseName).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func ReadOneRoom(idStr string) data.Room {
	client, ctx, cancel, err := Connect(dataBaseUri)
	collection := client.Database(databaseName).Collection("room")
	if err != nil {
		panic(err)
	}
	result := data.Room{}
	err = collection.FindOne(ctx, bson.M{"roomId": idStr}).Decode(&result)
	defer Close(client, ctx, cancel)
	return result
}

func ReadManyRooms(userId string) []data.Room {
	client, ctx, cancel, err := Connect(dataBaseUri)
	collection := client.Database(databaseName).Collection("room")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	results := make([]data.Room, 0)
	finalResults := make([]data.Room, 0)
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	for _, room := range results {
		if room.HostId == userId || contains(room.Members, userId) {
			finalResults = append(finalResults, room)
		}
	}
	defer Close(client, ctx, cancel)
	return finalResults
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ReadOneUser(idStr string) data.User {
	client, ctx, cancel, err := Connect(dataBaseUri)
	collection := client.Database(databaseName).Collection("user")
	if err != nil {
		panic(err)
	}
	result := data.User{}
	err = collection.FindOne(ctx, bson.M{"userId": idStr}).Decode(&result)
	defer Close(client, ctx, cancel)
	return result
}

func UpdateOne(col string, filter interface{}, update interface{}) int64 {
	client, ctx, cancel, err := Connect(dataBaseUri)
	if err != nil {
		panic(err)
	}
	collection := client.Database(databaseName).Collection(col)
	result, err := collection.UpdateOne(ctx, filter, update)
	defer Close(client, ctx, cancel)
	return result.ModifiedCount
}

func ReadAll(col string) {
	client, ctx, cancel, err := Connect(dataBaseUri)
	collection := client.Database(databaseName).Collection(col)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	defer Close(client, ctx, cancel)
	fmt.Printf("%+v\n", results) // Print with Variable Name
}
