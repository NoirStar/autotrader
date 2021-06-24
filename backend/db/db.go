package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/noirstar/autotrader/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// New makes new connection with mongodb
func New() *Connection {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test:" + utils.GetEnv("MONGO_DB_PASSWORD") + "@autotrader.byaqk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// CreateUser creates users
func CreateUser(client *mongo.Client) {
	collection := client.Database("autotrader").Collection("users")

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	/*
		Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var posts []Post
	if err = cur.All(ctx, &posts); err != nil {
		panic(err)
	}
	fmt.Println(posts)
}
