package service

import (
	"diplom/user-service/internal/models"
	"fmt"
)

func (s *Service) AddUser(req models.SaveUserRequest) error {
	studId, err := s.Store.CheckUser(req.StudentId)
	if err != nil {
		return err
	}

	if studId != "" {
		return fmt.Errorf("student id already exists")
	}

	return s.Store.SaveUser(req)
}
