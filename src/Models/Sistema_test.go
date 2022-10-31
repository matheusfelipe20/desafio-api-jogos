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

// Teste: Listar Campeonatos disponíveis (SUCESSO)
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

// Teste: Listar Jogos disponíveis (SUCESSO)
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
	resp, err := http.Get("http://localhost:8080/cpf/36806792979")
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

func TestNonExistentUserCPF(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/cpf/12345678910")
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
