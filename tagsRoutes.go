package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Tag struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId int    `json:"userId"`
	Type   int    `json:"type"`
}

func createTag(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	var tag Tag

	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, "Not a Valid Fail", http.StatusBadRequest)
		return
	}

	if _, err := db.Exec(context.Background(), "INSERT INTO security.tags (title, user_id) values($1,$2)", tag.Title, userId); err == nil {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
