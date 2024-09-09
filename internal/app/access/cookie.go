package access

import (
	"context"
	"errors"
	"github.com/KraDM09/gophermart/internal/app/util"
	"github.com/KraDM09/gophermart/internal/constants"
	"net/http"
)

type Cookie struct{}

func (c Cookie) Control(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token, err := r.Cookie(constants.CookieTokenKey)

		if errors.Is(err, http.ErrNoCookie) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userID := util.GetUserID(token.Value)

		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), constants.ContextUserIDKey, userID)
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
