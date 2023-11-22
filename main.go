package main

// underlying storage (in memory, on disk, s3, anything)
// server (http, tcp)

func main() {
	cfg := &Config{
		HTTPListenAddr: ":3000",
		WSListenAddr: ":4000",
		StoreProducerFunc: func() Storer {
			return NewMemoryStore()
		},
	}
	s, err := NewServer(cfg)
	if err != nil {
		panic(err)
	}
	
	
	s.Start()
}
