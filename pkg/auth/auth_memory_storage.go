package auth

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type mockUser struct {
}

func (u *mockUser) ID() string {
	return ""
}

func (u *mockUser) Username() string {
	return ""
}

func (u *mockUser) SetID(id string) {
}

type AuthMemoryStorage[T UserIdentifiable] struct {
	mtx   *sync.RWMutex
	users map[string]T
}

var _ AuthStorage[*mockUser] = (*AuthMemoryStorage[*mockUser])(nil)

func NewAuthMemoryStorage[T UserIdentifiable]() *AuthMemoryStorage[T] {
	return &AuthMemoryStorage[T]{
		mtx:   new(sync.RWMutex),
		users: make(map[string]T),
	}
}

func (s *AuthMemoryStorage[T]) Login(ctx context.Context, username string, password string) (T, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	for _, u := range s.users {
		if u.Username() == username {
			return u, nil
		}
	}
	var zero T
	return zero, ErrUserNotFound
}

func (s *AuthMemoryStorage[T]) Register(ctx context.Context, user T) (T, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	user.SetID(uuid.New().String())
	s.users[user.ID()] = user
	return user, nil
}
