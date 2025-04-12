package models

type SaveUserRequest struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	StudentId string `json:"student_id" db:"student_id"`
	Avatar    []byte `json:"avatar,omitempty" db:"avatar"`
}
