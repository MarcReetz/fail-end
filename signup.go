package main

import (
	"context"
	"encoding/json"
	"first-server/hash"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user string
	erro := db.QueryRow(context.Background(), "SELECT username FROM security.user WHERE username = $1", creds.Username).Scan(&user)
	switch erro {
	case nil:
		http.Error(w, "User exist already", http.StatusConflict)
		return
	case pgx.ErrNoRows:
		break
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var passwordHash string
	if passwordHash, err = hash.HashPassword(creds.Password); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if _, err := db.Exec(context.Background(), "insert into security.user (username,password) values($1,$2)", creds.Username, passwordHash); err == nil {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
