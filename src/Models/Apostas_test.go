package models

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

//Sucess
func Test_RealizarAposta(t *testing.T) {

	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 100,
  		"cliente_cpf": "368.067.929-79"
    }`)

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer(bilhete))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// testando se a venda foi realizada com sucesso
	expected := "Aposta realizada com sucesso!"
	respVend := &RespVenda{}

	if err := json.NewDecoder(resp.Body).Decode(respVend); err != nil {
		log.Fatal(err)
	}

	if respVend.Message != expected {
		t.Errorf("Expected %s, got %s", expected, respVend.Message)
	}

	if respVend.Code != "200" {
		t.Errorf("Expected %s, got %s", "200", respVend.Code)
	}
}

func Test_ErroValorLimite(t *testing.T) {

	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 200,
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
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	//Valor do Erro esperado em Bytes
	expected := `{"erro":"falha ao cadastrar, valor da aposta insuficiente, ou excedeu o limite do valor da aposta"}`

	eq, err := JSONBytesEqual(body, []byte(expected))
	if err != nil {
		log.Println(err)
	}

	//Comparação dos erros: erro recebido pelo body e o erro esperad
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, expected)
	}
}

func Test_ErroIDJogo(t *testing.T) {

	bilhete := []byte(`{
  		"id_jogo": 0,
  		"opcao_aposta": "1",
  		"valor_aposta": 100,
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
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	//Valor do Erro esperado em Bytes
	expected := `{"erro":"id do jogo é igual a 0"}`

	eq, err := JSONBytesEqual(body, []byte(expected))
	if err != nil {
		log.Println(err)
	}

	//Comparação dos erros: erro recebido pelo body e o erro esperad
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, expected)
	}
}

func Test_ErroClienteCpf(t *testing.T) {

	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 50,
  		"cliente_cpf": "400.067.929-82"
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
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	//Valor do Erro esperado em Bytes
	expected := `{"erro":"falha ao cadastrar, cpf inválido"}`

	eq, err := JSONBytesEqual(body, []byte(expected))
	if err != nil {
		log.Println(err)
	}

	//Comparação dos erros: erro recebido pelo body e o erro esperad
	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, expected)
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
