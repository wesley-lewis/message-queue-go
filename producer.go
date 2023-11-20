package main

import (
	"net/http"
	"strings"
	"fmt"

	"golang.org/x/exp/slog"
)

type Producer interface {
	Start() error
}

type HTTPProducer struct {
	listenAddr	string 
	producech	chan<- Message 
}

func NewHTTPProducer(listenAddr string, producech chan Message)*HTTPProducer {
	return &HTTPProducer{
		listenAddr: listenAddr,
		producech: producech,
	}
}

func(p *HTTPProducer) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	// commit 
	if r.Method == "GET" {
		
	}
	if r.Method == "POST" {
		if len(parts) != 2 {
			fmt.Println("invalid action")
			return
		}
		p.producech <- Message {
			Data: []byte("we don't know yet"),
			Topic: parts[1],
		}
	}
	fmt.Println(path)
}

func(p *HTTPProducer) handlePublish(w http.ResponseWriter, r *http.Request) {

}

func(p *HTTPProducer) Start() error {
	slog.Info("HTTP transport started", "port", p.listenAddr)
	return http.ListenAndServe(p.listenAddr, p)	
}
