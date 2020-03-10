package controllers

import (
	"html/template"
	"net/http"
)

type DataAtencion struct {
	Active string
}

// Atencion ...
func Atencion(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/atencion/atencion.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	atencion := &DataAtencion{
		Active: "atencion",
	}
	_ = t.ExecuteTemplate(w, "layout", atencion)
}