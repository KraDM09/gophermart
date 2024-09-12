package models

import "time"

type Withdrawal struct {
	Order        string    `json:"order"`
	Sum          float32   `json:"sum"`
	Processed_at time.Time `json:"processed_at"`
}
