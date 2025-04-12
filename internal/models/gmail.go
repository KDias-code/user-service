package models

type SendCodeRequest struct {
	Gmail string `json:"gmail"`
}

type CheckCodeRequest struct {
	Gmail string `json:"gmail"`
	Code  string `json:"code"`
}
