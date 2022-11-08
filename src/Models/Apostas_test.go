package models

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
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

//Criar bilhete de aposta (Erro Data Passada)
func TestCriarVendaErro_Data(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "São Paulo x Flamengo",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-30",
  			"opcao_aposta": "casa",
  			"valor_aposta": 150,
  			"cliente_nome": "Bello Moreira Alcântara",
  			"cliente_cpf": "659.102.554-52",
  			"cliente_nascimento": "01/01/2000"
    		}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se houve erro na data do evento
	expected := `{"erro":"falha ao cadastrar, insira a data do jogo, ou verfique se o jogo ainda está disponivel"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro CPF invalido)
func TestCriarVendaErro_CPF(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "São Paulo x Flamengo",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-09-30",
  			"opcao_aposta": "casa",
  			"valor_aposta": 150,
  			"cliente_nome": "Bello Moreira Alcântara",
  			"cliente_cpf": "709.102.554-82",
  			"cliente_nascimento": "01/01/2000"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se houve erro no CPF deve retornar uma mensagem de erro 
	var expected =  `{"erro":"falha ao cadastrar, cpf inválido"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro Menor de idade)
func TestCriarVendaErro_DataNascimento(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "São Paulo x Flamengo",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-09-30",
  			"opcao_aposta": "casa",
  			"valor_aposta": 150,
  			"cliente_nome": "Bello Moreira Alcântara",
  			"cliente_cpf": "659.102.554-52",
  			"cliente_nascimento": "01/01/2008"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se o usuário é menor de idade, deve retornar uma mensagem de erro
	var expected =  `{"erro":"falha ao cadastrar, usuário menor de idade"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro limite do valor da aposta excedido)
func TestCriarVendaErro_LimiteExcedido(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "São Paulo x Flamengo",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-09-30",
  			"opcao_aposta": "casa",
  			"valor_aposta": 550,
  			"cliente_nome": "Bello Moreira Alcântara",
  			"cliente_cpf": "659.102.554-52",
  			"cliente_nascimento": "01/01/2000"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se o valor da aposta é maior que o limite, deve retornar uma mensagem de erro
	var expected =  `{"erro":"falha ao cadastrar, valor da aposta insuficiente, ou excedeu o limite do valor da aposta"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro Nome do cliente vazio)
func TestCriarVendaErro_NomeCliente(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "São Paulo x Flamengo",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-09-30",
  			"opcao_aposta": "casa",
  			"valor_aposta": 150,
  			"cliente_nome": "",
  			"cliente_cpf": "659.102.554-52",
  			"cliente_nascimento": "01/01/2000"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	// testando se o nome do cliente está vazio, deve retornar uma mensagem de erro
	var expected =  `{"erro":"falha ao cadastrar, insira o nome do cliente"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro Nome do campeonato vazio)
func TestCriarVendaErro_NomeCampeonato(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "São Paulo x Flamengo",
  			"campeonato": "",
  			"data_jogo": "2022-09-30",
  			"opcao_aposta": "casa",
  			"valor_aposta": 125,
  			"cliente_nome": "Bello Moreira Alcântara",
  			"cliente_cpf": "659.102.554-52",
  			"cliente_nascimento": "01/01/2000"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se o nome do campeonato está vazio, deve retornar uma mensagem de erro
	var expected =  `{"erro":"falha ao cadastrar, insira o campeonato"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro Nome do titulo vazio)
func TestCriarVendaErro_NomeTitulo(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
  			"titulo_jogo": "",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-09-30",
  			"opcao_aposta": "1",
  			"valor_aposta": 125,
  			"cliente_nome": "Bello Moreira Alcântara",
  			"cliente_cpf": "659.102.554-52",
  			"cliente_nascimento": "01/01/2000"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se o nome do titulo está vazio, deve retornar uma mensagem de erro
	var expected =  `{"erro":"falha ao  cadastrar, insira o titulo do jogo"}`
	require.JSONEq(t, expected, string(body))
}

//Criar bilhete de aposta (Erro ID do jogo: 0)
func TestCriarVendaErro_IDjogo(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/vendas", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 0,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2023-10-31",
  			"opcao_aposta": "1",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "Helena",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "30/08/2001"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	
	// testando se o ID do jogo é 0, deve retornar uma mensagem de erro
	var expected =  `{"erro":"id do jogo é igual a 0"}`
	require.JSONEq(t, expected, string(body))
}
