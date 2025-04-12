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
