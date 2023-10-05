package session

import (
	"context"
	"sync"
)

type InMemoryStorage struct {
	mtx  *sync.RWMutex
	data map[string]any
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		mtx:  new(sync.RWMutex),
		data: make(map[string]any),
	}
}

func (s *InMemoryStorage) Read(ctx context.Context, key string) (any, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.data[key], nil
}

func (s *InMemoryStorage) Write(ctx context.Context, key string, value any) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.data[key] = value
	return nil
}
