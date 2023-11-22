package main

import (
	"log"
	"fmt"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _,  err := websocket.DefaultDialer.Dial("ws://localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
}
