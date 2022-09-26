package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Fail struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
	Hits        int    `json:"hits"`
	Tags        int    `json:"tags"`
}

type Hit struct {
	Id int `json:"id"`
}

const failIdKey contextKey = 1

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

func FailCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		failId := chi.URLParam(r, "failID")
		if failId == "" {
			http.Error(w, "No FailId suplied", http.StatusBadRequest)
			return
		}

		failIdInt, err := strconv.Atoi(failId)
		if err != nil {
			log.Println(err)
			http.Error(w, "No FailId suplied", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), failIdKey, failIdInt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func deletFail(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)

	var fail Fail

	if err := json.NewDecoder(r.Body).Decode(&fail); err != nil || fail.Id == 0 {
		http.Error(w, "Not A Valid Fail", http.StatusBadRequest)
		return
	}

	log.Println(fail.Id)

	if _, err := db.Exec(context.Background(), "DELETE FROM security.fail WHERE id = $1 AND user_id = $2", fail.Id, userId); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}

}

func getFail(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	failId := r.Context().Value(failIdKey)

	log.Println(reflect.TypeOf(userId))
	log.Println(reflect.TypeOf(failId))

	var fail Fail

	if err := db.QueryRow(context.Background(), "SELECT * FROM security.fail WHERE id = $1 AND user_id = $2", failId, userId).Scan(&fail.Id, &fail.Title, &fail.Description, &fail.UserId, &fail.Hits, &fail.Tags); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(fail)
	}

}

func addHit(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)

	var hit Hit
	if err := json.NewDecoder(r.Body).Decode(&hit); err != nil || hit.Id == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, err := db.Exec(context.Background(), "UPDATE security.fail SET hits = hits + 1 WHERE id = $1 AND user_id = $2 ", hit.Id, userId); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
