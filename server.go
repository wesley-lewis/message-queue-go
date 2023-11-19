package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	ListenAddr        string
	StoreProducerFunc StoreProducerFunc
}

type Server struct {
	Config *Config
	topics map[string]Storer
	consumers []Consumer
	producers []Producer
	quitch chan struct{}
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
		topics: make(map[string]Storer),
		quitch: make(chan struct{}),
	}, nil
}

func (s *Server) Start() {
	for _, consumer := range s.consumers {
		if err := consumer.Start(); err != nil {
			fmt.Println(err)
		}
	}

	for _, producer := range s.producers {
		if err := producer.Start(); err != nil {
			fmt.Println(err)
		}
	}
	<- s.quitch
	// http.ListenAndServe(s.Config.ListenAddr, s)
}

func (s *Server) createTopic(name string) bool {
	if _, ok := s.topics[name]; !ok {
		s.topics[name] = s.Config.StoreProducerFunc()
		return true
	}
	return false
}
