package models

type UserBalance struct {
	ID        int     `json:"id"`
	Balance   float64 `json:"balance"`
	Withdrawn int     `json:"withdrawn"`
}
