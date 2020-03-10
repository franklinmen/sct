package controllers

import (
	"html/template"
	"net/http"
)

type DataTurnos struct {
	Active string
}

// Turnos ...
func Turnos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/turnos/turnos.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	turnos := &DataTurnos{
		Active: "string",
	}
	_ = t.ExecuteTemplate(w, "layout", turnos)
}
