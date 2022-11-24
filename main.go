package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
	"github.com/matheusfelipe20/projeto-api-jogos/src/database"
	"github.com/matheusfelipe20/projeto-api-jogos/src/middleware"
)

func main() {
	database.Load()
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/campeonatos", controllers.ListarCampeonatos).Methods("GET")
	r.HandleFunc("/eventos", controllers.ListarEventos).Methods("GET")
	r.HandleFunc("/eventos/{id}", controllers.ListarEventosID).Methods("GET")
	r.HandleFunc("/eventos/campeonato/{id}", controllers.ListarEventosCampeonato).Methods("GET")
	r.HandleFunc("/eventos/data/{data}", controllers.ListarEventosData).Methods("GET")
	r.HandleFunc("/cpf/{cpf}", controllers.ListarUsuario).Methods("GET")
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/vendas", controllers.ListarVendas).Methods("GET")
	r.HandleFunc("/vendas", controllers.RealizarVenda).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", database.Porta), r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
