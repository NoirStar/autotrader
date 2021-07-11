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
	"golang.org/x/crypto/bcrypt"
)

var (
	errCreateUser     error = errors.New("DB Error : CreateUser")
	errLoginUser      error = errors.New("DB Error : LoginUser")
	errFindMarketData error = errors.New("DB Error : FindMarketData")
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
	collection := client.Database("common").Collection("user")

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

// LoginUser login users
func LoginUser(id, pw string) (*model.User, error) {
	client, ctx, cancel, err := New()
	if err != nil {
		return nil, err
	}
	collection := client.Database("common").Collection("user")
	var result *model.User

	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"id": id}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errLoginUser
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.PW), []byte(pw))
	if err != nil {
		return nil, errLoginUser
	}

	return result, nil
}

// CheckDuplicate checks duplication
func CheckDuplicate(params map[string]string) (bool, error) {
	client, ctx, cancel, err := New()
	if err != nil {
		return false, err
	}

	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database("common").Collection("user")

	for key, val := range params {
		filter := bson.M{key: val}
		num, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return false, err
		}

		if num > 0 {
			return true, nil
		}
	}
	return false, nil
}

// FindMarketData 디비 데이터 aggregate
func FindMarketData(minute time.Duration) (*model.Market, error) {
	time := (time.Now().Add(-minute)).Unix()
	client, ctx, cancel, err := New()
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database("coin").Collection("trade")

	matchAsk := bson.D{{"$match", bson.D{{
		"$and", bson.A{
			bson.D{{"ask_bid", "ASK"}},
			bson.D{{"trade_timestamp", bson.D{{"$gte", time}}}},
		},
	}}}}
	matchBid := bson.D{{"$match", bson.D{{
		"$and", bson.A{
			bson.D{{"ask_bid", "BID"}},
			bson.D{{"trade_timestamp", bson.D{{"$gte", time}}}},
		},
	}}}}

	groupAsk := bson.D{{
		"$group", bson.D{
			{"_id", "$code"},
			{"ask_count", bson.D{{"$sum", "$trade_volume"}}},
			{"ask_total", bson.D{{"$sum", bson.D{{"$multiply", bson.A{"$trade_volume", "$trade_price"}}}}}},
		},
	}}
	groupBid := bson.D{{
		"$group", bson.D{
			{"_id", "$code"},
			{"bid_count", bson.D{{"$sum", "$trade_volume"}}},
			{"bid_total", bson.D{{"$sum", bson.D{{"$multiply", bson.A{"$trade_volume", "$trade_price"}}}}}},
		},
	}}

	askCursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchAsk, groupAsk})
	if err != nil {
		return nil, err
	}
	bidCursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchBid, groupBid})
	if err != nil {
		return nil, err
	}
	var askData []bson.M
	var bidData []bson.M
	cAsk := make(chan bool)
	cBid := make(chan bool)
	market := make(model.Market)

	if err = askCursor.All(ctx, &askData); err != nil {
		return nil, errFindMarketData
	}
	if err = bidCursor.All(ctx, &bidData); err != nil {
		return nil, errFindMarketData
	}

	for _, val := range askData {
		code := fmt.Sprint(val["_id"])
		marketData := &model.MarketData{}
		market[code] = marketData
	}
	for _, val := range bidData {
		code := fmt.Sprint(val["_id"])
		marketData := &model.MarketData{}
		market[code] = marketData
	}

	go func() {
		for _, val := range askData {
			code := fmt.Sprint(val["_id"])
			marketData := &model.MarketData{}
			bsonData, err := bson.Marshal(val)
			if err != nil {
				cAsk <- false
				return
			}
			err = bson.Unmarshal(bsonData, &marketData)
			if err != nil {
				cAsk <- false
				return
			}
			market[code].AskCount = marketData.AskCount
			market[code].AskTotal = marketData.AskTotal
		}
		cAsk <- true
	}()
	go func() {
		for _, val := range bidData {
			code := fmt.Sprint(val["_id"])
			marketData := &model.MarketData{}
			bsonData, err := bson.Marshal(val)
			if err != nil {
				cBid <- false
				return
			}
			err = bson.Unmarshal(bsonData, &marketData)
			if err != nil {
				cBid <- false
				return
			}
			market[code].BidCount = marketData.BidCount
			market[code].BidTotal = marketData.BidTotal
		}
		cBid <- true
	}()

	return &market, nil
}
