package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/respostas"
)

// ListarEventos lista todos os eventos
// Chama a função GetJogos do services
func ListarEventos(w http.ResponseWriter, r *http.Request) {
	respostas.JSON(w, http.StatusOK, api.GetJogos())
}
