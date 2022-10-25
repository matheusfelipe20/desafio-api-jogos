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

//Teste: Listar Campeonatos disponíveis (SUCESSO)
func Test_ListarCampeonato(t *testing.T) {
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

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Teste: Listar Jogos disponíveis (SUCESSO)
func Test_ListarJogos(t *testing.T) {
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

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

func Test_UserCPF(t *testing.T) {

}

func Test_UserInvalidoCPF(t *testing.T) {

}

func Test_RealizarAposta(t *testing.T) {

}

func Test_UserMenorIdade(t *testing.T) {

}

func Test_JogoDataInvalida(t *testing.T) {

}

func Test_JogoCpfInvalido(t *testing.T) {

}

func Test_JogoInexistente(t *testing.T) {

}

func Test_CampeonatoInexistente(t *testing.T) {

}

func Test_LimeteExcedido(t *testing.T) {

}

func Test_ValorApostaNegativo(t *testing.T) {

}

func Test_ListarApostasRealizadas(t *testing.T) {

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
