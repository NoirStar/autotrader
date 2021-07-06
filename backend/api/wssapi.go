package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/noirstar/autotrader/db"
	"github.com/noirstar/autotrader/model"
	"github.com/noirstar/autotrader/utils"
)

var baseURI string = utils.GetEnv("UPBIT_WSS_BASE_URL")

// InitWSSClient 웹소켓 초기화
func InitWSSClient() {
	fmt.Println("Starting WSS Client")

	var limit int64 = 1024 * 1024 * 32
	var codes []string
	d := websocket.DefaultDialer
	cIncomingMsg := make(chan []byte)
	cSendingMsg := make(chan string)

	header := http.Header{}
	header.Set("origin", baseURI)

	ws, _, err := d.Dial(baseURI+"/websocket/v1", header)
	ws.SetReadLimit(limit)

	utils.CheckErr(err)

	defer ws.Close()

	info := []model.ResMarketCode{}
	if err := json.Unmarshal(GetMarketCode(), &info); err != nil {
		log.Fatalln(err)
	}

	go readWSMessage(ws, cIncomingMsg)
	go sendWSMessage(ws, cSendingMsg)

	for _, val := range info {
		codes = append(codes, val.Market)
	}
	a := model.NewReqForInfoWSS("trade", codes, true)
	cSendingMsg <- a.ReqForInfoJSON()

	// db 처리

	for msg := range cIncomingMsg {
		cMongo := make(chan []byte)

		go func() {
			defer close(cMongo)
			coinData := <-cMongo
			client, ctx, cancel, err := db.New()
			utils.CheckErr(err)
			collection := client.Database("autotrader").Collection("coins")

			defer client.Disconnect(ctx)
			defer cancel()
			data := model.Coin{}
			err = json.Unmarshal(coinData, &data)
			utils.CheckErr(err)
			_, err = collection.InsertOne(ctx, data)
			utils.CheckErr(err)
		}()

	}
}

func readWSMessage(ws *websocket.Conn, cIncomingMsg chan<- []byte) error {

	ws.SetPongHandler(func(appdata string) error {
		cIncomingMsg <- []byte(appdata)
		return nil
	})

	for {
		_, msg, err := ws.ReadMessage()
		utils.CheckErr(err)
		cIncomingMsg <- msg
	}
}

func sendWSMessage(ws *websocket.Conn, cSendingMsg chan string) error {
	for {
		select {
		case params := <-cSendingMsg:
			fmt.Println("ws send messages", params)
			err := ws.WriteMessage(websocket.TextMessage, []byte(params))
			utils.CheckErr(err)
		}
	}
}
