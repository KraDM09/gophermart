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

type RegisterRequest struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *UserHandler) RegisterHandler(
	ctx context.Context,
	rw http.ResponseWriter,
	r *http.Request,
) {
	var req RegisterRequest

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

	if user != nil {
		rw.WriteHeader(http.StatusConflict)
		json.NewEncoder(rw).Encode(map[string]string{"error": "User already exists"})
		return
	}

	password := req.Password + config.PasswordSalt

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	user, err = h.store.CreateUser(ctx, req.Login, string(hashedPassword))

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
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
