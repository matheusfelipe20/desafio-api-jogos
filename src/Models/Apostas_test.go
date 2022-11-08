package models

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_RealizarAposta(t *testing.T) {

	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 125,
  		"cliente_cpf": "368.067.929-79"
    }`)

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer(bilhete))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	// testando se a venda foi realizada com sucesso
	var expected = "Aposta realizada com sucesso!"
	if string(body) != expected {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}
