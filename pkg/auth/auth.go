package auth

import (
	"context"

	"github.com/mrparano1d/archs-go/pkg/session"
)

type AuthManager[T UserIdentifiable] struct {
	storage AuthStorage[T]
	sess    *session.SessionManager
}

func NewAuthManager[T UserIdentifiable](storage AuthStorage[T], sessionManager *session.SessionManager) *AuthManager[T] {
	return &AuthManager[T]{
		storage: storage,
		sess:    sessionManager,
	}
}

func (m *AuthManager[T]) Login(ctx context.Context, username string, password string) (*session.Session, error) {
	user, err := m.storage.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}

	sess, err := m.sess.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (m *AuthManager[T]) Register(ctx context.Context, user T) (T, error) {
	return m.storage.Register(ctx, user)
}
