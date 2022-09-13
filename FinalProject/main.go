package main

import (
	"finalProject/app"
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// func main() {
// 	srv := server.Server{Repository: &server.AccountRepositoryInMemory{}}

// 	r := chi.NewRouter()

// 	r.Use(middleware.Logger)
// 	r.Post("/NewAccount", srv.HandleAccount)
// 	// r.Get("/tweets", srv.Repository.Tweets)

// 	// http.HandleFunc("/tweet", srv.Handletweet)
// 	log.Fatal(http.ListenAndServe(":5001", r))
// }

func main() {

	r := chi.NewRouter()
	r.Post("/account/create", app.HandleCreateAccount)
	log.Fatal(http.ListenAndServe(":5001", r))
}
