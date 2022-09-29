package main

import (
	"context"
	"encoding/json"
	"first-server/client/utils"
	"first-server/pointifyer"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Tag struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId int    `json:"userId"`
	//Type   int    `json:"type"`
}

var changeAllowedTagField = []string{"title"}

const tagIdKey contextKey = 2

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

func tagCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tagId := chi.URLParam(r, "tagID")
		if tagId == "" {
			http.Error(w, "No TagId suplied", http.StatusBadRequest)
			return
		}

		failIdInt, err := strconv.Atoi(tagId)
		if err != nil {
			log.Println(err)
			http.Error(w, "No FailId suplied", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), tagIdKey, failIdInt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllTags(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	log.Println("hit")

	var tag []Tag

	if rows, err := db.Query(context.Background(), "SELECT * FROM security.tags WHERE user_id = $1", userId); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		for rows.Next() {
			var temp Tag
			columns, _ := pointifyer.Pointify(&temp)
			rows.Scan(columns...)
			tag = append(tag, temp)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(tag)
	}
}

func getTag(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	tagNumber := r.Context().Value(tagIdKey)

	tag := Tag{}

	columns, _ := pointifyer.Pointify(&tag)

	if err := db.QueryRow(context.Background(), "SELECT * FROM security.tags WHERE user_id = $1 AND id = $2", userId, tagNumber).Scan(columns...); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(tag)
	}
}

func deleteTag(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	tagNumber := r.Context().Value(tagIdKey)

	if _, err := db.Exec(context.Background(), "DELETE FROM security.tags WHERE user_id = $1 AND id = $2", userId, tagNumber); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func updateTag(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	tagId := r.Context().Value(tagIdKey)

	var result map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	var keys string
	var elements []any
	var num = 3

	for key, element := range result {
		if utils.Contains(changeAllowedTagField, key) {
			keys += key + " = $" + strconv.Itoa(num) + " ,"
			num++
			elements = append(elements, element)
		}
	}

	keys = strings.TrimRight(keys, ",")

	SQLString := "UPDATE security.tags SET " + keys + " WHERE id = $1 AND user_id = $2"

	log.Println(SQLString)

	if _, err := db.Exec(context.Background(), SQLString, append([]any{tagId, userId}, elements...)...); err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
