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

func (s *Service) UpdateUser(req models.SaveUserRequest) error {
	return s.Store.UpdateUser(req)
}

func (s *Service) GetUser(studId string) (models.SaveUserRequest, error) {
	stud, err := s.Store.GetUser(studId)
	if err != nil {
		return models.SaveUserRequest{}, err
	}

	return stud, nil
}
