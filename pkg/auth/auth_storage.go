package auth

import (
	"context"
	"errors"
)

type UserIdentifiable interface {
	ID() string
	Username() string
	SetID(id string)
}

type AuthStorage[U UserIdentifiable] interface {
	Login(ctx context.Context, username string, password string) (U, error)
	Register(ctx context.Context, user U) (U, error)
}

var ErrUserNotFound = errors.New("user not found")
