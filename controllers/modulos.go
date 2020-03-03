package controllers

import (
	"html/template"
	"net/http"
)

// Modulos ...
func Modulos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/modulos/modulos.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	active := "modulos"
	_ = t.ExecuteTemplate(w, "layout", active)
}
