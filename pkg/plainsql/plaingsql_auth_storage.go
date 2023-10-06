package plainsql

import (
	"context"
	"database/sql"

	"github.com/mrparano1d/archs-go/pkg/auth"
)

type User struct {
	id       string
	username string
	password string
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetID(id string) {
	u.id = id
}

type SqlAuthStorage struct {
	client *sql.DB
}

var _ auth.AuthStorage[*User] = &SqlAuthStorage{}

func NewSqlAuthStorage(client *sql.DB) *SqlAuthStorage {
	return &SqlAuthStorage{client: client}
}

func (s *SqlAuthStorage) Login(ctx context.Context, username string, password string) (*User, error) {
	var user User

	err := s.client.QueryRowContext(ctx, "SELECT id, username, password FROM goarch.users WHERE username = $1 AND password = $2", username, password).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *SqlAuthStorage) Register(ctx context.Context, user *User) (*User, error) {
	_, err := s.client.ExecContext(ctx, "INSERT INTO goarch.users (username, password) VALUES ($1, $2)", user.username, user.password)
	if err != nil {
		return nil, err
	}

	user, err = s.Login(ctx, user.username, user.password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
