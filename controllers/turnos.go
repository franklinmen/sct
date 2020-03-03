package controllers

import (
	"html/template"
	"net/http"
)

// Turnos ...
func Turnos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/turnos/turnos.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	active := "turnos"
	_ = t.ExecuteTemplate(w, "layout", active)
}
