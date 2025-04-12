package models

type SaveUserRequest struct {
	Name      string `json:"name" db:"name"`
	StudentId string `json:"student_id" db:"student_id"`
	Avatar    string `json:"avatar,omitempty" db:"avatar"`
}
