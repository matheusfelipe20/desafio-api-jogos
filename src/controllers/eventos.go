package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/respostas"
)

// ListarEventos lista todos os eventos
// Chama a função GetJogos do services
func ListarEventos(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, api.GetJogos())
}

func ListarEventosID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	intkey, _ := strconv.Atoi(key)

	for _, jogo := range api.GetJogos() {
		if jogo.ID == intkey {
			json.NewEncoder(w).Encode(jogo)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func ListarEventosData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, jogo := range api.GetJogos() {
		if jogo.Data == key {
			json.NewEncoder(w).Encode(jogo)
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func ListarEventosCampeonato(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	intkey, _ := strconv.Atoi(key)

	for _, campeonato := range api.GetCampeonatos() {
		if campeonato.ID == intkey {
			json.NewEncoder(w).Encode(campeonato)
		}
	}

	for _, jogo := range api.GetJogos() {
		if jogo.ID_Campeonato == intkey {
			json.NewEncoder(w).Encode(jogo)
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
