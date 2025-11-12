package api

import (
	"github.com/Just1a2Noob/accouting-journal-app/internal/database"
)

type LoggedUser struct {
	name  string
	token string
}

type ApiConfig struct {
	Database database.Queries
	Secret   string
	User     LoggedUser
}
