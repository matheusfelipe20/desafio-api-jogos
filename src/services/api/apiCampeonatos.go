package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/matheusfelipe20/projeto-api-jogos/src/Models"
)

// GetJogos irá consumir a api do heroku https://desafio-api-jogos.herokuapp.com/campeonatos
// Transforma o corpo da resposta em um tipo campeonato
func GetCampeonatos() []models.Campeonato {

	api := "https://desafio-api-jogos.herokuapp.com/campeonatos"
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
