package constants

import "time"

type contextKey string

const (
	ContextUserIDKey contextKey = "user_id"
	CookieTokenKey   string     = "token"
	Lifetime                    = 24 * 7 * time.Hour
)

const (
	LoyaltyOrderStatusRegistered string = "REGISTERED"
	LoyaltyOrderStatusInvalid    string = "INVALID"
	LoyaltyOrderStatusProcessing string = "PROCESSING"
	LoyaltyOrderStatusProcessed  string = "PROCESSED"
)
