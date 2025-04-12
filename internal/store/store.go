package store

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type IStore interface {
	SaveCode(ctx context.Context, code, gmail string) error
	CheckCode(ctx context.Context, gmail string) (string, error)
}
type Store struct {
	redis    *redis.Client
	postgres *sqlx.DB
}

func NewStore(redis *redis.Client, postgres *sqlx.DB) *Store {
	return &Store{
		redis:    redis,
		postgres: postgres,
	}
}
