package controllers

import (
	"log"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
)

// ListarEventos lista todos os eventos
// Chama a função GetJogos do services
func ListarEventos(w http.ResponseWriter, r *http.Request) {
	log.Println(api.GetJogos())
}
