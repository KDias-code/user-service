package service

import (
	"context"
	"diplom/user-service/internal/models"
	"diplom/user-service/internal/store"
	"diplom/user-service/pkg/gmail"
)

type IService interface {
	SendCode(toGmail string) error
	CheckCode(ctx context.Context, gmail, userCode string) (bool, error)
	AddUser(req models.SaveUserRequest) error
	UpdateUser(req models.SaveUserRequest) error
	GetUser(studId string) (models.SaveUserRequest, error)
}
type Service struct {
	Gmail gmail.IGmail
	Store store.IStore
}

func NewService(gmail gmail.IGmail, store store.IStore) *Service {
	return &Service{
		Gmail: gmail,
		Store: store,
	}
}
