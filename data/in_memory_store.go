package data

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
)

type InMemoryStore struct {
	points map[string]int
	mutex  sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		points: make(map[string]int),
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

func (s *InMemoryStore) CreatePointsEntry(points int) (id string, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id, err = s.createId()
	if err != nil {
		return "", err
	}

	s.points[id] = points

	return id, nil
}

func (s *InMemoryStore) createId() (id string, err error) {
	uuid := make([]byte, 16)
	_, err = rand.Read(uuid)
	if err != nil {
		return "", err
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	var buf [36]byte
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])

	return string(buf[:]), nil
}
