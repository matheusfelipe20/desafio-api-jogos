package funcoes

import "testing"

type limiteCheck struct {
	limite   float64
	valorApostado float64
	esperado bool
}

func TestValidarLimite(t *testing.T){

	limite := []limiteCheck{
		{
			limite: 100.0,
			valorApostado: 50.0,
			esperado: true,
		},
		{
			limite: 100.0,
			valorApostado: 150.0,
			esperado: false,
		},
		{
			limite: 100.0,
			valorApostado: 100.0,
			esperado: true,
		},
	}

	for _, lim := range limite {
		resultado := ValidadeLimiteValor(lim.limite, lim.valorApostado)
		if resultado != lim.esperado {
			t.Errorf("Resultado esperado: %v, resultado obtido: %v", lim.esperado, resultado)
		}
	}

}

	
