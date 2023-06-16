package main

import (
	"net/http"

	inter "exemplo.com/internal"
	"github.com/go-chi/chi/v5"
)

func main() {
	base := inter.NewBase()
	handers := inter.NewBaseHandler(*base)
	r := chi.NewRouter()
	r.Get("/", handers.GetHandler)
	r.Post("/", handers.PostHandler)
	http.ListenAndServe(":8000", r)
}
