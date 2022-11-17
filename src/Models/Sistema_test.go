package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_LocalhostOPen(t *testing.T) {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected %d, received %d", http.StatusAccepted, resp.StatusCode)
	}
}

//Teste para consultar todos os campeonatos disponiveis
func TestGetCampeonato(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/campeonatos")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	//Valor esperado pela consulta
	campeonatosCadastrado := []byte(`[{"id":30,"titulo":"Brasileirão - Serie A"},{"id":35,"titulo":"Copa América - Feminina"},{"id":36,"titulo":"Uruguai - Primeira Divisão"}]`)

	eq, err := JSONBytesEqual(body, campeonatosCadastrado)
	if err != nil {
		log.Println(err)
	}

	//Comparação do valor esperado com recebido
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, campeonatosCadastrado)
	}

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
}

//Teste para verficar a filtragem por ID Campeonato
func TestGetJogoIDCampeonato(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos/campeonato/30")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Teste para verificar o error de consulta por ID Campeonato | ID Campeonato não encontrado (NotFound)
func TestNonExistentJogoIDCampeonato(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos/campeonato/20")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Teste para verficar a filtragem por ID Jogo
func TestGetJogoID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos/354858757161276")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	//Valor esperado pela busca por ID Jogo
	jogoCadastradoID := `{"id":354858757161276,"titulo":"Ceará x Avaí","id_campeonato":30,"data":"2022-08-20","opcoes":[{"1":10.14},{"x":2.5},{"2":1.7}],"limites":[{"1":650},{"x":750},{"2":500}]}`

	eq, err := JSONBytesEqual(body, []byte(jogoCadastradoID))
	if err != nil {
		log.Println(err)
	}
	//Comparação do valor esperado com recebido
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, jogoCadastradoID)
	}

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

}

//Teste para verificar o error de consulta por ID Jogo | ID Jogo não encontrado (NotFound)
func TestNonExistentJogo(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos/65420162165499")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Teste para verificar a filtragem de eventos(jogos) por data
func TestGetJogoData(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos/data/2022-11-11")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	//Valor esperado pela consulta
	jogoCadastradoData := []byte(`{"id":65489162165499,"titulo":"Deportivo Maldonado x Torque da Cidade de Montevideu","id_campeonato":36,"data":"2022-11-11","opcoes":[{"1":1.25},{"x":4.5},{"2":3.9}],"limites":[{"1":0},{"x":0},{"2":0}]}`)

	eq, err := JSONBytesEqual(body, jogoCadastradoData)
	if err != nil {
		log.Println(err)
	}
	//Comparar valor esperado com o recebido
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, jogoCadastradoData)
	}

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

}

//Teste para verificar o error de consulta por Data | Data não encontrada (NotFound)
func TestNonExistentDataJogo(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos/data/2022-10-01")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}
