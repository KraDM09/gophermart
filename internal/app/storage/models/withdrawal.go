package models

import "time"

type Withdrawal struct {
	Number       string    `json:"number"`
	Sum          int       `json:"sum"`
	Processed_at time.Time `json:"processed_at"`
}
