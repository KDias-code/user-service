package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
)

func (s *Service) SendCode(toGmail string) error {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return err
	}
	code := fmt.Sprintf("%06d", n.Int64())

	err = s.Gmail.SendCode(toGmail, code)
	if err != nil {
		return err
	}

	err = s.Store.SaveCode(context.Background(), code, toGmail)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CheckCode(ctx context.Context, gmail, userCode string) (bool, error) {
	code, err := s.Store.CheckCode(ctx, gmail)
	if err != nil {
		return false, err
	}

	if code != userCode {
		return false, nil
	}

	return true, nil
}
