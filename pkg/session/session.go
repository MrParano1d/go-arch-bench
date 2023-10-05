package session

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID       string `json:"id"`
	User     any 	`json:"user"`
	ExpireAt time.Time `json:"expire_at"`
}

type SessionManager struct {
	storage SessionStorage
}

func NewSessionManager(storage SessionStorage) *SessionManager {
	return &SessionManager{storage}
}

func (m *SessionManager) Create(ctx context.Context, user any) (*Session, error) {
	s := &Session{
		ID:       uuid.New().String(),
		User:     user,
		ExpireAt: time.Now().Add(24 * time.Hour),
	}
	err := m.storage.Write(ctx, s.ID, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *SessionManager) Get(ctx context.Context, id string) (*Session, error) {
	s, err := m.storage.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.(*Session), nil
}
