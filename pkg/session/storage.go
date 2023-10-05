package session

import "context"

type SessionStorage interface {
	Read(ctx context.Context, key string) (any, error)
	Write(ctx context.Context, key string, value any) error
}
