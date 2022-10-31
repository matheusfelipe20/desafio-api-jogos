package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("Not port")
	}
	fmt.Printf("Listening Port %s\n", port)

	r := mux.NewRouter()

	r.HandleFunc("/campeonatos", controllers.ListarCampeonatos).Methods("GET")
	r.HandleFunc("/eventos", controllers.ListarEventos).Methods("GET")
	r.HandleFunc("/cpf", controllers.ListarUsuario).Methods("GET")
	r.HandleFunc("/", home).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
