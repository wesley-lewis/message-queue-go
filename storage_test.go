package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	s := NewMemoryStore()
	var err error
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("foobarbaz_%d", i)
		_, err = s.Push([]byte(key))
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 100; i++ {
		data, err := s.Fetch(int32(i))
		if err != nil {
			panic(err)
		}

		value := fmt.Sprintf("foobarbaz_%d", i)
		assert.Equal(t, data, []byte(value))
	}
}
