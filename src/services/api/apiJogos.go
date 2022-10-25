package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/matheusfelipe20/projeto-api-jogos/src/Models"
)

// GetJogos irá consumir a api do heroku https://desafio-api-jogos.herokuapp.com/jogos
// Transforma o corpo da resposta em um tipo jogo
func GetJogos() []models.Campeonato {

	api := "https://desafio-api-jogos.herokuapp.com/jogos"
	response, err := http.Get(api)
	if err != nil {
		log.Println("erro na resposta")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	var campeonatos []models.Campeonato
	json.Unmarshal(responseData, &campeonatos)
	return campeonatos
}
