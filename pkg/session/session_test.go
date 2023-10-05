package session_test

import (
	"context"
	"testing"

	"github.com/mrparano1d/archs-go/pkg/session"
)

type testUser struct {
	ID   int
	Name string
}

func TestSessionManager_Create(t *testing.T) {
	manager := session.NewSessionManager(session.NewInMemoryStorage())
	ctx := context.Background()

	user := &testUser{ID: 1, Name: "test"}

	s, err := manager.Create(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	if s.ID == "" {
		t.Fatal("empty session id")
	}

	if s.User != user {
		t.Fatal("invalid user")
	}

	if s.ExpireAt.IsZero() {
		t.Fatal("invalid expire at")
	}
}

func TestSessionManager_Get(t *testing.T) {
	manager := session.NewSessionManager(session.NewInMemoryStorage())
	ctx := context.Background()

	user := &testUser{ID: 1, Name: "test"}

	s, err := manager.Create(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	s2, err := manager.Get(ctx, s.ID)
	if err != nil {
		t.Fatal(err)
	}

	if s2.ID != s.ID {
		t.Fatal("invalid session id")
	}

	if s2.User != s.User {
		t.Fatal("invalid user")
	}

	if !s2.ExpireAt.Equal(s.ExpireAt) {
		t.Fatal("invalid expire at")
	}
}
