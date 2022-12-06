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

var url = "http://localhost:8080"

// Sucess
func Test_RealizarAposta(t *testing.T) {

	//Cadastro do bilhete de aposta
	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 100,
  		"cliente_cpf": "368.067.929-79"
    }`)

	resp, err := http.Post(url+"/vendas", "application/json",
		bytes.NewBuffer(bilhete))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	//Testando se a venda foi realizada com sucesso
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

//test para validar o limite do valor de aposta
func Test_ErroValorLimite(t *testing.T) {

	//cadastro do bilhete
	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 200,
  		"cliente_cpf": "368.067.929-79"
    }`)
	//O limite para o valor da aposta desse jogo "354858757161272" é 150

	resp, err := http.Post(url+"/vendas", "application/json",
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
	expected := `{"erro":"StatusCode: 400, Erro: Limite do valor da aposta excedido}`

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

	//Cadastro do bilhete de aposta
	bilhete := []byte(`{
  		"id_jogo": 0,
  		"opcao_aposta": "1",
  		"valor_aposta": 100,
  		"cliente_cpf": "368.067.929-79"
    }`)
	//ID "0" Incorreto

	resp, err := http.Post(url+"/vendas", "application/json",
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
	expected := `{"erro":"StatusCode: 400, Erro: ID de jogo não encontrado}`

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

	//Cadastro do bilhete de aposta
	bilhete := []byte(`{
  		"id_jogo": 354858757161272,
  		"opcao_aposta": "1",
  		"valor_aposta": 50,
  		"cliente_cpf": "400.067.929-82"
    }`)
	//CPF do cliente é inválido (não existente)

	resp, err := http.Post(url+"/vendas", "application/json",
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
	expected := `{"erro":"StatusCode: 400, Erro: CPF inválido}`

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
