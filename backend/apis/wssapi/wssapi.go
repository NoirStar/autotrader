package wssapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/env"
	"github.com/noirstar/autotrading/backend/utils/money"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

var baseURI string = env.GetEnv("UPBIT_WSS_BASE_URL")

// InitWSSClient 웹소켓 초기화
func InitWSSClient() {
	fmt.Println("Starting WSS Client")

	var limit int64 = 1024 * 1024 * 32
	d := websocket.DefaultDialer
	cIncomingMsg := make(chan []byte)
	cSendingMsg := make(chan string)

	header := http.Header{}
	header.Set("origin", baseURI)

	ws, _, err := d.Dial(baseURI+"/websocket/v1", header)
	ws.SetReadLimit(limit)

	myerr.CheckErr(err)

	defer ws.Close()

	go readWSMessage(ws, cIncomingMsg)
	go sendWSMessage(ws, cSendingMsg)

	cd := []string{"KRW-DOGE"}
	a := models.NewReqForInfo("trade", cd, true)
	cSendingMsg <- a.ReqForInfoJSON()

	idx := 0
	for {

		msg := <-cIncomingMsg
		data := models.ResTrade{}
		err := json.Unmarshal(msg, &data)
		myerr.CheckErr(err)

		m, err := money.NewMoney(data.TradePrice)
		myerr.CheckErr(err)

		fmt.Println(idx, "Received:", m.Display())
		idx++
	}

}

func readWSMessage(ws *websocket.Conn, cIncomingMsg chan<- []byte) error {

	ws.SetPongHandler(func(appdata string) error {
		cIncomingMsg <- []byte(appdata)
		return nil
	})

	for {
		_, msg, err := ws.ReadMessage()
		myerr.CheckErr(err)
		cIncomingMsg <- msg
	}
}

func sendWSMessage(ws *websocket.Conn, cSendingMsg chan string) error {
	for {
		select {
		case params := <-cSendingMsg:
			fmt.Println("ws send messages", params)
			err := ws.WriteMessage(websocket.TextMessage, []byte(params))
			myerr.CheckErr(err)
		}
	}
}
