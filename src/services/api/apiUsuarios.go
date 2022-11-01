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

	apiCpf1 := ("https://apijogos.herokuapp.com/cpf/36806792979")
	response, _ := http.Get(apiCpf1)

	apiCpf2 := ("https://apijogos.herokuapp.com/cpf/36806792979")
	response2, _ := http.Get(apiCpf2)

	apiCpf3 := ("https://apijogos.herokuapp.com/cpf/36806792979")
	response3, _ := http.Get(apiCpf3)

	apiCpf4 := ("https://apijogos.herokuapp.com/cpf/36806792979")
	response4, _ := http.Get(apiCpf4)

	responseData1, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	responseData2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	responseData3, err := ioutil.ReadAll(response3.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	responseData4, err := ioutil.ReadAll(response4.Body)
	if err != nil {
		log.Println("erro ao converter o json")
	}

	usr1 := models.Usuario{}
	usr2 := models.Usuario{}
	usr3 := models.Usuario{}
	usr4 := models.Usuario{}

	var usuarios []models.Usuario
	json.Unmarshal(responseData1, &usr1)
	json.Unmarshal(responseData2, &usr2)
	json.Unmarshal(responseData3, &usr3)
	json.Unmarshal(responseData4, &usr4)

	usuarios = append(usuarios, usr1)
	usuarios = append(usuarios, usr2)
	usuarios = append(usuarios, usr3)
	usuarios = append(usuarios, usr4)

	return usuarios
}
