package store

import (
	"context"
	"time"
)

func (s *Store) SaveCode(ctx context.Context, code, gmail string) error {
	err := s.redis.Set(ctx, gmail, code, 2*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CheckCode(ctx context.Context, gmail string) (string, error) {
	codeDb, err := s.redis.Get(ctx, gmail).Result()
	if err != nil {
		return "", err
	}

	return codeDb, nil
}
