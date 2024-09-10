package models

import "time"

type Order struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Number     string    `json:"number"`
	UploadedAt time.Time `json:"uploaded_at"`
	Status     string    `json:"status"`
	Accrual    *float64  `json:"accrual"`
}
