package main

import (
	"errors"
	"sync"
)

type InMemoryStore struct {
	receipts map[string]Receipt
	points   map[string]int
	mutex    sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		receipts: make(map[string]Receipt),
		points:   make(map[string]int),
	}
}

func (s *InMemoryStore) GetPoints(receiptId string) (int, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	points, exists := s.points[receiptId]
	if !exists {
		return 0, errors.New("receipt not found")
	}

	return points, nil
}
