package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"highroller-go/src/data"
	"log"
	"time"
)

var dataBaseUri = "mongodb://localhost:27017"
var databaseName = "highroller"

func Update() {

}

func Delete() {

}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func Ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
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
	fmt.Println("ID:" + idStr)
	//docID, err := primitive.ObjectIDFromHex(idStr)
	result := data.Room{}
	err = collection.FindOne(ctx, bson.M{"roomId": idStr}).Decode(&result)
	defer Close(client, ctx, cancel)
	fmt.Printf("%+v\n", result) // Print with Variable Name
	return result
}

func ReadManyRooms(userId string) []data.Room {
	client, ctx, cancel, err := Connect(dataBaseUri)
	collection := client.Database(databaseName).Collection("room")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var results []data.Room
	var finalResults []data.Room
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	for _, room := range results {
		if room.HostId == userId || contains(room.Members, userId) {
			finalResults = append(finalResults, room)
		}
	}
	defer Close(client, ctx, cancel)
	fmt.Printf("%+v\n", results) // Print with Variable Name
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
	fmt.Println("ID:" + idStr)
	//docID, err := primitive.ObjectIDFromHex(idStr)
	result := data.User{}
	err = collection.FindOne(ctx, bson.M{"userId": idStr}).Decode(&result)
	defer Close(client, ctx, cancel)
	fmt.Printf("%+v\n", result) // Print with Variable Name
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
	return result.UpsertedCount
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

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}
