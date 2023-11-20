package main

import (
	"fmt"

	"golang.org/x/exp/slog"
)

type Message struct {
	Topic string 
	Data []byte
}

type Config struct {
	ListenAddr        string
	StoreProducerFunc StoreProducerFunc
}

type Server struct {
	Config *Config
	topics map[string]Storer
	consumers []Consumer
	producers []Producer
	producech chan Message
	quitch chan struct{}
}

func NewServer(cfg *Config) (*Server, error) {
	producech := make(chan Message)
	return &Server{
		Config: cfg,
		topics: make(map[string]Storer),
		quitch: make(chan struct{}),
		producech: producech,
		producers: []Producer{NewHTTPProducer(cfg.ListenAddr, producech)},
	}, nil
}

func (s *Server) Start() {
	// for _, consumer := range s.consumers {
	// 	if err := consumer.Start(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

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
			if err := s.publish(msg); err != nil {
				slog.Error("failed to publish", err)
			}
		}
	}
}

func(s *Server) publish(msg Message) error {
	s.createTopicIfNotExist(msg.Topic)

	return nil
}

func (s *Server) createTopicIfNotExist(topic string)  {
	if _, ok := s.topics[topic]; !ok {
		s.topics[topic] = s.Config.StoreProducerFunc()
		slog.Info("created new topic", "topic", topic)
	}
}
