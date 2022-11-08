package models

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
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

//Teste para consultar todos os jogos
func TestGetEventos(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/eventos")
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
	jogosCadastrados := []byte(`[{"id":354858757161272,"titulo":"São Paulo x Flamengo","id_campeonato":30,"opcoes":[{"1":2.5},{"x":3.1},{"2":1.5}]},
	{"id":354858757161273,"titulo":"Fluminense x Palmeiras","id_campeonato":30,"opcoes":[{"1":1.25},{"x":4.5},{"2":3.9}]},
	{"id":354858757161274,"titulo":"Botafogo x Santos","id_campeonato":30,"opcoes":[{"1":10.14},{"x":2.5},{"2":1.7}]},
	{"id":354858757161275,"titulo":"Vasco x Atlético","id_campeonato":30,"opcoes":[{"1":1.25},{"x":4.5},{"2":3.9}]},
	{"id":354858757161276,"titulo":"Ceará x Avaí","id_campeonato":30,"opcoes":[{"1":10.14},{"x":2.5},{"2":1.7}]},
	{"id":354858324654689,"titulo":"Colômbia x Chile","id_campeonato":35,"opcoes":[{"1":1.63},{"x":3.84},{"2":5.09}]},
	{"id":354858324654690,"titulo":"Equador x Paraguai","id_campeonato":35,"opcoes":[{"1":5.77},{"x":4.32},{"2":1.5}]},
	{"id":65489162165498,"titulo":"Liverpool FC x AlbionFC","id_campeonato":36,"opcoes":[{"1":5.77},{"x":4.32},{"2":1.5}]},
	{"id":65489162165499,"titulo":"Deportivo Maldonado x Torque da Cidade de Montevideu","id_campeonato":36,"opcoes":[{"1":1.25},{"x":4.5},{"2":3.9}]}]`)

	eq, err := JSONBytesEqual(body, jogosCadastrados)
	if err != nil {
		log.Println(err)
	}

	//Comparação do valor esperado com recebido
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, jogosCadastrados)
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
	jogoCadastradoID := []byte(`{"id":354858757161276,"titulo":"Ceará x Avaí","id_campeonato":30,"opcoes":[{"1":10.14},{"x":2.5},{"2":1.7}]}`)

	eq, err := JSONBytesEqual(body, jogoCadastradoID)
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

//Funções para comparar os Json's

// JSONEqual comparando dois Json
func JSONEqual(a, b io.Reader) (bool, error) {
	var j, j2 interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// JSONBytesEqual compara o JSON em fatias de dois bytes.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
