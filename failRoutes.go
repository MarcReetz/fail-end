package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Fail struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
	Hits        int    `json:"hits"`
	Tags        int    `json:"tags"`
}

func createFail(w http.ResponseWriter, r *http.Request) {
	var fail Fail

	if err := json.NewDecoder(r.Body).Decode(&fail); err != nil {
		http.Error(w, "Not a Valid Fail", http.StatusBadRequest)
		return
	}

	if fail.Title == "" {
		http.Error(w, "No Title Provided", http.StatusBadGateway)
		return
	}

	userId := r.Context().Value(authenticatedUserKey)

	if _, err := db.Exec(context.Background(), "INSERT INTO security.fail (title, description, user_id) values($1,$2,$3)", fail.Title, fail.Description, userId); err == nil {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func deletFail(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)

	var fail Fail

	if err := json.NewDecoder(r.Body).Decode(&fail); err != nil && fail.Id != 0 {
		http.Error(w, "Not A Valid Fail", http.StatusBadRequest)
		return
	}

	log.Println(fail.Id)

	if _, err := db.Exec(context.Background(), "DELETE FROM security.fail WHERE id = $1 AND user_id = $2", fail.Id, userId); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
