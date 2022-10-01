package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

const dbNoRowsError string = "no rows in result set"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	log.Println(os.Getenv("DATABASE_URL"))
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	log.Println(poolConfig.ConnConfig.Database)

	db, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/", func(r chi.Router) {
		r.Post("/api/login", Login)
		r.Post("/api/signup", signup)
		//Private
		r.Route("/api/user", func(r chi.Router) {
			r.Use(CheckToken)
			r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("login worked"))
			})
			r.Post("/fail", createFail)
			r.Get("/fail", getAllFails)
			r.Route("/fail/{failID}", func(r chi.Router) {
				r.Use(FailCtx)
				r.Get("/", getFail)
				r.Put("/", updateFail)
				r.Delete("/", deletFail)
				r.Put("/hit", addHit)
			})

			r.Post("/tag", createTag)
			r.Get("/tag", getAllTags)
			r.Route("/tag/{tagID}", func(r chi.Router) {
				r.Use(tagCtx)
				r.Delete("/", deleteTag)
				r.Get("/", getTag)
				r.Put("/", updateTag)
			})
		})
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":3000", r)
}
