package main

import(
	"github.com/gorilla/websocket"
)

type Consumer interface {
	Start() error
}

func foo() {
	websocket.DefaultDialer.Dial("ws:/foo", nil)
}
