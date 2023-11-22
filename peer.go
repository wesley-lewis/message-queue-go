package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
)

type Peer interface {
	Send([]byte) error
}

type WSPeer struct {
	conn *websocket.Conn
}

func NewWSPPeer(conn *websocket.Conn) Peer {
	p :=  &WSPeer {
		conn: conn,
	}
	go p.readLoop()
	return p
}

func(p *WSPeer) readLoop() {
	var msg WSMessage 

	for {
		if err := p.conn.ReadJSON(&msg); err != nil {
			slog.Error("WS peer readJSON error", "error", err)
			return 
		}
		if err := p.handleMessage(msg); err != nil {
			slog.Error("WS peer read error", "error", err)
			return 
		}
	}
}

func (p *WSPeer) handleMessage(msg WSMessage) error {
	fmt.Printf("handling msg -> %+v\n", msg)
	return nil
}

func(p *WSPeer) Send(b []byte) (error) {
	return p.conn.WriteMessage(websocket.BinaryMessage, b)
}
