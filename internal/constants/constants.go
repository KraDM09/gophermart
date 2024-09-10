package constants

import "time"

type contextKey string

const (
	ContextUserIDKey contextKey = "user_id"
	CookieTokenKey   string     = "token"
	Lifetime                    = 24 * 7 * time.Hour
)

const (
	LOYALTY_ORDER_STATUS_REGISTERED string = "REGISTERED"
	LOYALTY_ORDER_STATUS_INVALID           = "INVALID"
	LOYALTY_ORDER_STATUS_PROCESSING        = "PROCESSING"
	LOYALTY_ORDER_STATUS_PROCESSED         = "PROCESSED"
)
