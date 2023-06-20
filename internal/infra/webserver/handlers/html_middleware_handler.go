package handlers

import (
	"net/http"
	"text/template"
)

func MiddlewareHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/middleware.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
