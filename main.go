package main

import (
	"fmt"
	"github.com/devc94/sct/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r.HandleFunc("/", controllers.HomePage)
	r.HandleFunc("/atencion", controllers.Atencion)
	r.HandleFunc("/modulos", controllers.Modulos)
	r.HandleFunc("/tramites", controllers.Tramites)
	r.HandleFunc("/turnos", controllers.Turnos)

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}