package handlers

import (
	"context"
	"encoding/json"
	"github.com/KraDM09/gophermart/internal/app/config"
	"github.com/KraDM09/gophermart/internal/app/util"
	"github.com/KraDM09/gophermart/internal/constants"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginRequest struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *UserHandler) LoginHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
		return
	}

	user, err := h.store.GetUserByLogin(ctx, req.Login)

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		rw.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(rw).Encode(map[string]string{"error": "Invalid login or password"})
		return
	}

	password := req.Password + config.PasswordSalt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(rw).Encode(map[string]string{"error": "Invalid login or password"})
		return
	}

	token, err := util.GenerateJWT(user.ID)

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(rw, &http.Cookie{
		Name:    constants.CookieTokenKey,
		Value:   token,
		Expires: time.Now().Add(constants.Lifetime),
		Path:    "/",
	})

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
