package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// ResTicker 현재가 응답 데이터 포멧
type ResTicker struct {
	Type               string  `json:"type"`
	Code               string  `json:"code"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	Change             float64 `json:"change"` // RISE : 상승 EVEN : 보합 FALL : 하락
	ChangePrice        float64 `json:"change_price"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
	TradeDate          string  `json:"trade_date"` // yyyyMMdd
	TradeTime          string  `json:"trade_time"` // HHmmss
	TradeTimestamp     int32   `json:"trade_timestamp"`
	AskBid             string  `json:"ask_bid"` // ASK : 매도 BID : 매수
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"` // yyyy-MM-dd
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"` // yyyy-MM-dd
	TradeStatus        string  `json:"trade_status"`
	MarketState        string  `json:"market_state"` // PREVIEW : 입금지원 ACTIVE : 거래지원가능 DELISTED : 거래지원종료
	MarketStateForIos  string  `json:"market_state_for_ios"`
	IsTradingSuspended bool    `json:"is_trading_suspended"`
	DelistingDate      string  `json:"delisting_date"`
	MarketWarning      string  `json:"market_warning"` // NONE : 해당없음 CAUTION : 투자유의
	Timestamp          int32   `json:"timestamp"`
	StreamType         string  `json:"stream_type"` // SNAPSHOT : 스냅샷 REALTIME : 실시간
}

var baseURI string = "wss://api.upbit.com/websocket/v1"

func main() {
	initWSSClient()

}

func initWSSClient() {
	fmt.Println("Starting Client")

	var limit int64 = 1024 * 1024 * 32
	d := websocket.DefaultDialer
	cIncomingMsg := make(chan []byte)
	cSendingMsg := make(chan string)
	header := http.Header{}
	origin := strings.Split(baseURI, "/websocket")[0]

	header.Set("origin", origin)
	ws, _, err := d.Dial(baseURI, nil)
	ws.SetReadLimit(limit)

	checkErr(err, "Dial Error")

	defer ws.Close()

	go readWSMessage(ws, cIncomingMsg)
	go sendWSMessage(ws, cSendingMsg)

	cSendingMsg <- `[{"ticket":"UNIQUE_TICKET"},{"type":"ticker","codes":["KRW-BTC"]}]`

	for {
		msg := <-cIncomingMsg

		fmt.Println("Message Received:", string(msg))
	}

}

func readWSMessage(ws *websocket.Conn, cIncomingMsg chan<- []byte) error {

	ws.SetPongHandler(func(appdata string) error {
		cIncomingMsg <- []byte(appdata)
		return nil
	})

	for {
		msgType, message, err := ws.ReadMessage()
		checkErr(err, "Read Message Error")
		fmt.Println("Message Type:", msgType)
		cIncomingMsg <- message
	}
}

func sendWSMessage(ws *websocket.Conn, cSendingMsg chan string) error {
	for {
		select {
		case params := <-cSendingMsg:
			fmt.Println("ws send messages", params)
			err := ws.WriteMessage(websocket.TextMessage, []byte(params))
			checkErr(err, "Write Message Error")
		}
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
