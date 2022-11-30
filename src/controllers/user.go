package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/Models/funcoes"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/respostas"
)

// ListarUsuario irá buscar um usuário através do CPF
func ListarUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cpf := vars["cpf"]

	for _, usuario := range api.GetUser() {
		if funcoes.Formated(usuario.Cpf) == cpf {
			respostas.JSON(w, http.StatusOK, usuario)
			return
		}
	}

	respostas.JSON(w, http.StatusNotFound, "Usuário não encontrado")
}
