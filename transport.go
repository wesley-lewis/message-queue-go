package main 

import (
	"net/http"
	"fmt"
	"strings"

	"golang.org/x/exp/slog"
)

type Consumer interface {
	Start() error
}

type Producer interface {
	Start() error
}

type HTTPProducer struct {
	listenAddr string 
}

func NewHTTPProducer(listenAddr string) *HTTPProducer {
	return &HTTPProducer{
		listenAddr: listenAddr,
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
	}
	fmt.Println(path)
}

func(p *HTTPProducer) handlePublish(w http.ResponseWriter, r *http.Request) {

}

func(p *HTTPProducer) Start() error {
	slog.Info("HTTP transport started", "port", p.listenAddr)
	return http.ListenAndServe(p.listenAddr, p)	
}
