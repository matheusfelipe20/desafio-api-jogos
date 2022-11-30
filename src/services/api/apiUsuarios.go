package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/matheusfelipe20/projeto-api-jogos/src/Models"
)

// GetUser irá consumir a api do heroku https://apijogos.herokuapp.com/cpf
func GetUser() []models.Usuario {

	cpfs := []string{"83604761794", "36806792979", "46110334499", "23130011480"}
	url := "https://apijogos.herokuapp.com/cpf/"

	var usrs []models.Usuario
	var usr models.Usuario

	for _, cpf := range cpfs{
		request := fmt.Sprintf("%s%s", url, cpf)
		response, _ := http.Get(request)

		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseBody, &usr)
		usrs = append(usrs, usr)

	}

	return usrs
}
