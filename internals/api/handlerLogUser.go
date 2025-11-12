package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Just1a2Noob/accouting-journal-app/internals/auth"
)

func (cfg *ApiConfig) HandlerLogUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	login := struct {
		username string
		password string
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&login)
	if err != nil {
		ErrorResponse(w, "Error decoding json body", err, http.StatusBadRequest)
		return
	}

	user, err := cfg.Database.SearchUser(context.Background(), login.username)
	if err != nil {
		ErrorResponse(w, "Cannot find user", err, http.StatusBadRequest)
		return
	}

	err = auth.CheckPasswordHash(login.password, user.HashPass)
	if err != nil {
		ErrorResponse(w, "Password is invalid", err, http.StatusBadRequest)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.Secret, auth.ExpiresIn)
	if err != nil {
		ErrorResponse(w, "Cannot create token", err, http.StatusInternalServerError)
		return
	}
	_, err = auth.ValidateJWT(token, cfg.Secret)
	if err != nil {
		ErrorResponse(w, "Invalid JWT", err, http.StatusInternalServerError)
		return
	}

	cfg.User.name = user.Username
	cfg.User.token = token

	w.WriteHeader(http.StatusNoContent)
}
