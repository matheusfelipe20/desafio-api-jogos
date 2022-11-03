package funcoes

import "testing"

type dataCheck struct {
	data     string
	esperado bool
}

// teste para verificar se o usuario é maior de idade
func Test_UsuarioMaiorDeIdade(t *testing.T) {

	usuarios := []dataCheck{
		{
			data:     "20/11/2000",
			esperado: true,
		}, // testando se o usuario é maior de idade
		{
			data:     "17/11/2010",
			esperado: false,
		}, // testando usuario com menos de 18 anos
		{
			data:     "00/00/0000",
			esperado: false,
		}, // testando usuario com data de nascimento invalida
	}

	for _, usr := range usuarios {
		resultado := ValidadeDataNascimento(usr.data)
		if resultado != usr.esperado {
			t.Errorf("Resultado esperado: %v, resultado obtido: %v", usr.esperado, resultado)
		}
	}

}

// teste para verificar se a venda passou da validade
func Test_DataVendaValidade(t *testing.T) {

	vendas := []dataCheck{

		{
			data:     "2022-11-01",
			esperado: true,
		}, // testando a data dias antes do evento
		{
			data:     "2022-11-05",
			esperado: true,
		}, // testando a data um dia antes do jogo como sendo válida
		{
			data:     "2022-11-06",
			esperado: false,
		}, // testando a data no dia do jogo sendo uma data inválida
		{
			data:     "2022-11-07",
			esperado: false,
		}, // testando a data um dia depois do jogo sendo uma data inválida
		{
			data:     "",
			esperado: false,
		}, // testando se a data estiver vazia
		{
			data:     "0000-00-00",
			esperado: false,
		}, // testando se a data estiver com formato inválido
	}

	for _, venda := range vendas {
		resultado := ValidadeDataVenda(venda.data)
		if resultado != venda.esperado {
			t.Errorf("Resultado esperado: %v, resultado obtido: %v", venda.esperado, resultado)
		}
	}

}
