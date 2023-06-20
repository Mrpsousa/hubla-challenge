package handlers

import (
	"net/http"
	"text/template"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
