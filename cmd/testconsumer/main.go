package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _,  err := websocket.DefaultDialer.Dial("ws://localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}

	msg := WSMessage {
		Action: "subscribe",
		Topics: []string{"foobarbar"},
	}
	
	conn.WriteJSON(msg)
}

type WSMessage struct {
	Action	string		`json:"action"`
	Topics	[]string	`json:"topics"`
}
