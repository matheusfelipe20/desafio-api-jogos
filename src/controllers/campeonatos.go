package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/respostas"
)

// ListarCampeonatos lista todos os campeonatos
// Chama a função GetCampeonatos do services
func ListarCampeonatos(w http.ResponseWriter, r *http.Request) {
	respostas.JSON(w, http.StatusOK, api.GetCampeonatos())
}
