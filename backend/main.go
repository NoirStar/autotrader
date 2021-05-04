package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/websocket"
)

var baseURL string = "wss://api.upbit.com/websocket/v1"

func main() {
	initWebsocketClient()
}

func initWebsocketClient() {
	fmt.Println("Starting Client")
	cIncomingMsg := make(chan []byte)
	ws, err := websocket.Dial(baseURL, "", strings.Split(baseURL, "/websocket")[0])
	checkErr(err)

	_, err = ws.Write([]byte("PING"))
	checkErr(err)

	go readClientMessages(ws, cIncomingMsg)

	for {
		msg := <-cIncomingMsg
		fmt.Println("Message Received:", msg)
	}

}

func readClientMessages(ws *websocket.Conn, cIncomingMsg chan<- []byte) {
	for {
		var msg = make([]byte, 512)
		//err := websocket.Message.Receive(ws, &msg)
		n, err := ws.Read(msg)
		checkErr(err)
		cIncomingMsg <- msg[:n]
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
