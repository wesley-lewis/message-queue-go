package main

import (
	"fmt"
	"sync"

	"golang.org/x/exp/slog"
)

type Message struct {
	Topic string 
	Data []byte
}

type Config struct {
	HTTPListenAddr			string
	WSListenAddr			string 
	StoreProducerFunc		StoreProducerFunc
}

type Server struct {
	Config		*Config

	mu			sync.RWMutex
	peers		map[Peer] bool

	topics		map[string]Storer
	consumers []Consumer
	producers []Producer
	producech chan Message
	quitch chan struct{}
}

func NewServer(cfg *Config) (*Server, error) {
	producech := make(chan Message)
	s :=  &Server{
		Config: cfg,
		topics: make(map[string]Storer),
		quitch: make(chan struct{}),
		peers: make(map[Peer]bool),
		producech: producech,
		producers: []Producer{NewHTTPProducer(cfg.HTTPListenAddr, producech)},
	}

	s.consumers = []Consumer{NewWSConsumer(s.Config.WSListenAddr, s)}
	return s, nil
}

func (s *Server) Start() {
	for _, consumer := range s.consumers {
		go func(c Consumer) {
			if err := c.Start(); err != nil {
				fmt.Println(err)
			}
		}(consumer)
	}

	for _, producer := range s.producers {
		go func(p Producer) {
			if err := p.Start(); err != nil {
				fmt.Println(err)
			}
		}(producer)
	}
	s.loop()
	// http.ListenAndServe(s.Config.ListenAddr, s)
}

func(s *Server) loop() {
	for {
		select {
		case <- s.quitch:
			return
		case msg := <- s.producech:
			fmt.Println("Produced: ", msg)
			offset, err := s.publish(msg) 
			if err != nil {
				slog.Error("failed to publish", err)
			}else {
				slog.Info("published to", "offset", offset)
			}
		}
	}
}

func(s *Server) publish(msg Message) (int32, error) {
	store := s.getStoreForTopic(msg.Topic)
	return store.Push(msg.Data)
}

func (s *Server) getStoreForTopic(topic string) Storer {
	if _, ok := s.topics[topic]; !ok {
		s.topics[topic] = s.Config.StoreProducerFunc()
		slog.Info("created new topic", "topic", topic)
	}
	
	return s.topics[topic]
}

func (s *Server) AddConn(p Peer) {
	s.mu.Lock()
	defer s.mu.Unlock() 
	slog.Info("added new peers", "peer", p)
	s.peers[p] = true
}
