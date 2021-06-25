package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/noirstar/autotrader/model"
	"github.com/noirstar/autotrader/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	errCreateUser error = errors.New("DB Error : CreateUser")
)

// New makes new connection with mongodb
func New() (*mongo.Client, context.Context, context.CancelFunc, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test:" + utils.GetEnv("MONGO_DB_PASSWORD") + "@autotrader.byaqk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		return nil, nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		client.Disconnect(ctx)
		cancel()
		return nil, nil, nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		client.Disconnect(ctx)
		cancel()
		return nil, nil, nil, err
	}
	return client, ctx, cancel, nil
}

// CreateUser creates users
func CreateUser(user *model.User) error {
	client, ctx, cancel, err := New()
	if err != nil {
		return err
	}
	collection := client.Database("autotrader").Collection("users")

	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			bson.M{"id": user.ID},
			bson.M{"email": user.Email},
			bson.M{"nickname": user.Nickname},
		},
	}

	num, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	if num == 0 {
		res, err := collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	}
	return errCreateUser
}
