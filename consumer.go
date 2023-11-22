package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
)

var upgrader = websocket.Upgrader{}

type Consumer interface {
	Start() error
}

type WSConsumer struct {
	ListenAddr string 
	server *Server
}

func NewWSConsumer(listenAddr string, s *Server ) *WSConsumer {
	return &WSConsumer{
		ListenAddr: listenAddr,
		server: s,
	}
}

func (ws *WSConsumer) Start() error {
	slog.Info("websocket consumer started", "port", ws.ListenAddr)
	http.ListenAndServe(":4000", ws)
	return nil
}

func(ws *WSConsumer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return 
	}

	p := NewWSPPeer(conn)
	ws.server.AddConn(p)
	fmt.Println(conn)
}

func foo() {
	websocket.DefaultDialer.Dial("ws:/foo", nil)
}
