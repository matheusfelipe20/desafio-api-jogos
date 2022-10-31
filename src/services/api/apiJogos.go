package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/matheusfelipe20/projeto-api-jogos/src/Models"
)

func GetJogos() []models.Jogo {

	api := "https://apijogos.herokuapp.com/jogos"
	response, err := http.Get(api)
	if err != nil {
		log.Println("erro na resposta")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	var jogos []models.Jogo
	json.Unmarshal(responseData, &jogos)
	return jogos
}
