package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/noirstar/autotrader/model"
	"github.com/noirstar/autotrader/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection mongo db connection struct
type Connection struct {
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
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
	return &Connection{
		Client: client,
		Ctx:    ctx,
		Cancel: cancel,
	}
}

// CreateUser creates users
func (c *Connection) CreateUser(user *model.User) {
	collection := c.Client.Database("autotrader").Collection("users")

	defer c.Client.Disconnect(c.Ctx)
	defer c.Cancel()

	res, err := collection.InsertOne(c.Ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
