package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcReetz/fail-end/client/utils"
	"github.com/MarcReetz/fail-end/pointifyer"

	"github.com/go-chi/chi/v5"
)

type Fail struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
	Hits        int    `json:"hits"`
	Tags        []int  `json:"tags"`
}

var changeAllowedFailField = []string{"title", "description", "tags"}

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
	failId := r.Context().Value(failIdKey)

	if _, err := db.Exec(context.Background(), "DELETE FROM security.fail WHERE id = $1 AND user_id = $2", failId, userId); err != nil {
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

	fail := Fail{}

	columns, _ := pointifyer.Pointify(&fail)

	err := db.QueryRow(context.Background(), "SELECT * FROM security.fail WHERE id = $1 AND user_id = $2", failId, userId).Scan(columns...)
	log.Println(fail)
	if err != nil {
		log.Println(err)
		if err.Error() == dbNoRowsError {
			http.Error(w, "No Such Fail", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(fail)
	}
}

func addHit(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	failId := r.Context().Value(failIdKey)

	if _, err := db.Exec(context.Background(), "UPDATE security.fail SET hits = hits + 1 WHERE id = $1 AND user_id = $2 ", failId, userId); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func updateFail(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	failId := r.Context().Value(failIdKey)

	var result map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	var keys string
	var elements []any
	var num = 3

	for key, element := range result {
		if utils.Contains(changeAllowedFailField, key) {
			keys += key + " = $" + strconv.Itoa(num) + " ,"
			num++
			elements = append(elements, element)
		}
	}

	keys = strings.TrimRight(keys, ",")

	SQLString := "UPDATE security.fail SET " + keys + " WHERE id = $1 AND user_id = $2"

	log.Println(SQLString)

	if _, err := db.Exec(context.Background(), SQLString, append([]any{failId, userId}, elements...)...); err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}

}

func getAllFails(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(authenticatedUserKey)
	var fails []Fail
	if rows, err := db.Query(context.Background(), "SELECT * FROM security.fail WHERE user_id = $1", userId); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		for rows.Next() {
			var temp Fail
			columns, _ := pointifyer.Pointify(&temp)
			rows.Scan(columns...)
			fails = append(fails, temp)
		}
		if rows.Err() != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(fails)
	}
}
