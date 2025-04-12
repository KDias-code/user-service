package store

import (
	"diplom/user-service/internal/models"
	"fmt"
)

func (s *Store) CheckUser(studentId string) (string, error) {
	var studId string
	query := fmt.Sprintf(`SELECT student_id FROM users WHERE student_id = '%s'`, studentId)

	err := s.postgres.Get(&studId, query)
	if err != nil {
		return "", err
	}

	return studId, nil
}

func (s *Store) SaveUser(req models.SaveUserRequest) error {
	query := fmt.Sprintf(`INSERT INTO users (name, student_id) VALUES (:name, :student_id)`)

	_, err := s.postgres.NamedExec(query, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(req models.SaveUserRequest) error {
	var query string
	if req.Avatar != nil && req.Name != "" {
		query = fmt.Sprintf(`UPDATE users SET name = '%s', avatar = '%v' WHERE student_id = '%s'`, req.Name, req.Avatar, req.StudentId)
	} else if req.Avatar != nil && req.Name == "" {
		query = fmt.Sprintf(`UPDATE users SET avatar = '%v' WHERE student_id = '%s'`, req.Avatar, req.StudentId)
	} else if req.Name != "" && req.Avatar == nil {
		query = fmt.Sprintf(`UPDATE users SET name = '%s' WHERE student_id = '%s'`, req.Name, req.StudentId)
	}

	_, err := s.postgres.NamedExec(query, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUser(studId string) (models.SaveUserRequest, error) {
	var user = models.SaveUserRequest{}
	query := fmt.Sprintf(`SELECT * FROM users WHERE student_id = '%s'`, studId)

	err := s.postgres.Get(&user, query)
	if err != nil {
		return user, err
	}

	return user, nil
}
