package controllers

import (
	"log"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
)

// ListarCampeonatos lista todos os campeonatos
// Chama a função GetCampeonatos do services
func ListarCampeonatos(w http.ResponseWriter, r *http.Request) {
	log.Println(api.GetCampeonatos())
}
