package controllers

import (
	"html/template"
	"net/http"
)

// Tramites ...
func Tramites(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/tramites/tramites.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	active := "tramites"
	_ = t.ExecuteTemplate(w, "layout", active)
}
