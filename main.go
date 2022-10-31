package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/campeonatos", controllers.ListarCampeonatos).Methods("GET")
	r.HandleFunc("/eventos", controllers.ListarEventos).Methods("GET")
	r.HandleFunc("/cpf", controllers.ListarUsuario).Methods("GET")
	r.HandleFunc("/", home).Methods("GET")

	port := map[bool]string{true: os.Getenv("PORT"), //Se minha variavel PORT não for declarada
		false: "8080"}[os.Getenv("PORT") != ""] //Assumo que a porta é 8080
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
