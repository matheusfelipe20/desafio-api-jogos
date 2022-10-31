package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/respostas"
)

// ListarUsuario irá buscar um usuário através do CPF
func ListarUsuario(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, api.GetUser())
}
