package main

import (
	"finalProject/app"
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	r := chi.NewRouter()
	r.Post("/account/create", app.HandleCreateAccount)
	log.Fatal(http.ListenAndServe(":5001", r))
}
