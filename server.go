package main

type Config struct {
	ListenAddr string
	Store      Storer
}

type Server struct {
	Config *Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
	}, nil
}
