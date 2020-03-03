package main

import (
	"fmt"
	"github.com/devc94/sct/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomePage)

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
