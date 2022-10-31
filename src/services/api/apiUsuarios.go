package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/matheusfelipe20/projeto-api-jogos/src/Models"
)

// GetUser irá consumir a api do heroku https://apijogos.herokuapp.com/cpf
func GetUser() []models.Usuario {

	api := ("https://apijogos.herokuapp.com/cpf/36806792979")
	response, err := http.Get(api)
	if err != nil {
		log.Println("erro na resposta")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	var usuarios []models.Usuario
	json.Unmarshal(responseData, &usuarios)
	return usuarios
}
