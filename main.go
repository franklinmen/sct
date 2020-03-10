package main

import (
	"fmt"
	"github.com/devc94/sct/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/olahol/melody.v1"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	m := melody.New()
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r.HandleFunc("/", controllers.HomePage)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	}).Methods("GET")

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.HandleFunc("/atencion", controllers.Atencion)

	r.HandleFunc("/modulos", controllers.Modulos)

	r.HandleFunc("/tramites", controllers.Tramites)
	r.HandleFunc("/tramites/lista", controllers.ListaTramites)
	r.HandleFunc("/tramites/nuevo", controllers.NuevoTramite).Methods("GET")
	r.HandleFunc("/tramites/nuevo", controllers.NuevoTramite).Methods("POST")
	r.HandleFunc("/tramites/eliminar/{id}", controllers.EliminarTramite).Methods("DELETE")

	r.HandleFunc("/turnos", controllers.Turnos)

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}