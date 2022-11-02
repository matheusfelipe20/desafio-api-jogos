package funcoes

import (
	"reflect"
	"testing"
)

type oddsOpcoes struct {
	odd      string
	oddsList []map[string]float64
	esperado map[string]float64
}

type oddsLimites struct {
	odd      string
	oddsList []map[string]float64
	esperado float64
}

// irá comparar a opção de aposta selecionada pelo jogador com as opções de aposta disponíveis e retornar a opção e o valor da opção
func Test_CompararOpcaoDeAposta(t *testing.T) {

	odds := []oddsOpcoes{
		{
			odd:      "x",
			oddsList: []map[string]float64{{"1": 2.5}, {"x": 3.5}, {"2": 4.5}},
			esperado: map[string]float64{"x": 3.5},
		},
		{
			odd:      "1",
			oddsList: []map[string]float64{{"1": 5.0}, {"x": 2.5}, {"2": 4.1}},
			esperado: map[string]float64{"1": 5.0},
		},
		{
			odd:      "2",
			oddsList: []map[string]float64{{"1": 2.0}, {"x": 1.5}, {"2": 2.1}},
			esperado: map[string]float64{"2": 2.1},
		},
	}

	var resultado map[string]float64

	for _, odd := range odds {
		resultString, resultNum := CompararOpcaoAposta(odd.odd, odd.oddsList)
		resultado = map[string]float64{resultString: resultNum}
		if !reflect.DeepEqual(resultado, odd.esperado) {
			t.Errorf("Resultado esperado: %v, resultado obtido: %v", odd.esperado, resultado)
		}
	}
}

// irá comparar o limite com a opção de aposta selecionada pelo jogador e retornar o valor do limite referente a opção
func Test_CompararLimiteComOpcao(t *testing.T) {

	odds := []oddsLimites{
		{
			odd:      "x",
			oddsList: []map[string]float64{{"1": 150.0}, {"x": 500.0}, {"2": 750.0}}, // limites
			esperado: 500.0,
		},
		{
			odd:      "1",
			oddsList: []map[string]float64{{"1": 100.0}, {"x": 200.0}, {"2": 300.0}}, // limites
			esperado: 100.0,
		},
		{
			odd:      "2",
			oddsList: []map[string]float64{{"1": 50.0}, {"x": 100.0}, {"2": 150.0}}, // limites
			esperado: 150.0,
		},
	}

	var resultado float64
	for _, odd := range odds {
		resultado = CompararLimiteComOpcao(odd.odd, odd.oddsList)
		if resultado != odd.esperado {
			t.Errorf("Resultado esperado: %v, resultado obtido: %v", odd.esperado, resultado)
		}
	}

}
