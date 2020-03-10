package controllers

import (
	"html/template"
	"net/http"
)

type DataHome struct {
	Active string
}

// HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/home.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	home := &DataHome{
		Active: "home",
	}
	_ = t.ExecuteTemplate(w, "layout", home)
}
