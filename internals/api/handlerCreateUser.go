package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Just1a2Noob/accouting-journal-app/internal/database"
	"github.com/Just1a2Noob/accouting-journal-app/internals/auth"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userdb UserDB
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userdb)
	if err != nil {
		ErrorResponse(w, "Error decoding json body", err, http.StatusBadRequest)
		return
	}

	hash_pass, err := auth.HashPassword(userdb.Password)
	if err != nil {
		ErrorResponse(w, "Error hashing password", err, http.StatusInternalServerError)
		return
	}
	err = auth.CheckPasswordHash(userdb.Password, hash_pass)
	if err != nil {
		ErrorResponse(w, "Error password does not match hash", err, http.StatusInternalServerError)
		return
	}

	user, err := cfg.Database.CreateUsers(context.Background(), database.CreateUsersParams{
		ID:        uuid.New(),
		Username:  userdb.Username,
		HashPass:  hash_pass,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Role:      userdb.Role,
		Section:   userdb.Section,
		Email:     userdb.Email,
	})
	if err != nil {
		ErrorResponse(w, "Error creating user in database", err, http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Section:  user.Section,
		Roles:    user.Role,
	}
	data, err := json.Marshal(response)
	if err != nil {
		ErrorResponse(w, "Error giving response to client", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
