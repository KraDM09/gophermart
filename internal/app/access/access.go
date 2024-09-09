package access

import (
	"net/http"
)

type Access interface {
	Control(next http.Handler) http.Handler
}
