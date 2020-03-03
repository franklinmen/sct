package controllers

import (
	"html/template"
	"net/http"
)

// HomePage...
func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/views/home.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = t.ExecuteTemplate(w, "layout", nil)
}