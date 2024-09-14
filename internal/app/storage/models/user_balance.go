package models

type UserBalance struct {
	ID        int     `json:"id"`
	Balance   float32 `json:"balance"`
	Withdrawn float32 `json:"withdrawn"`
}
