package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/devc94/sct/config"
	"github.com/devc94/sct/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type DataTramites struct {
	T      []models.Tramite
	Active string
}

func getTramites(query string) (tramites []models.Tramite, err error) {
	db := config.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := models.Tramite{}
		err = rows.Scan(&t.ID, &t.Nombre, &t.Iniciales, &t.TerceraEdad)
		if err != nil {
			return nil, err
		}
		tramites = append(tramites, t)
	}
	return tramites, nil
}

// ListaTramites...
func ListaTramites(w http.ResponseWriter, r *http.Request) {
	t, err := getTramites("SELECT * FROM tramites;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(t) > 0 {
		err = json.NewEncoder(w).Encode(t)
	} else {
		trm := &models.Tramite{}
		err = json.NewEncoder(w).Encode(trm)
	}
}

// Tramites ...
func Tramites(w http.ResponseWriter, r *http.Request) {
	tr, err := getTramites("SELECT * FROM tramites;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tramites := &DataTramites{
		T:      tr,
		Active: "tramites",
	}
	t, err := template.New("").Funcs(template.FuncMap{
		"add": func(number int) int {
			return number + 1
		},
	}).ParseFiles("public/views/tramites/tramites.html", "public/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(tramites)
	_ = t.ExecuteTemplate(w, "layout", tramites)
}

// NuevoTramite ...
func NuevoTramite(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("public/views/tramites/nuevo-tramite.html", "public/views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tramites := &DataTramites{
			T:      nil,
			Active: "nuevo-tramite",
		}
		_ = t.ExecuteTemplate(w, "layout", tramites)
	case "POST":
		t := models.Tramite{}

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db := config.GetConnection()
		defer db.Close()
		st, err := db.Prepare("INSERT INTO tramites VALUES(?, ?, ?);")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		st.Exec(t.Nombre, t.Iniciales, t.TerceraEdad)
	}
}

// EliminarTramite ...
func EliminarTramite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idTramite := params["id"]
	db := config.GetConnection()
	defer db.Close()
	st, err := db.Prepare("DELETE FROM tramites WHERE id = ?;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	st.Exec(idTramite)
}
