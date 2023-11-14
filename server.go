package main

type Config struct {
	ListenAddr string
}

type Server struct {
	Store  Storer
	Config *Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
	}, nil
}
