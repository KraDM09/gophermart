package models

import "time"

type Withdrawal struct {
	Order        string    `json:"order"`
	Sum          int       `json:"sum"`
	Processed_at time.Time `json:"processed_at"`
}
