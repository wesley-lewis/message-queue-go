package main

import (
	"fmt"
	"sync"
)

type Storer interface {
	Push([]byte) error
	Fetch(int32) ([]byte, error)
}

type MemoryStore struct {
	mu   sync.RWMutex
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) (int32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, b)
	return int32(len(s.data) - 1), nil
}

func (s *MemoryStore) Fetch(offset int32) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if int32(len(s.data)) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}

	return s.data[offset], nil
}
