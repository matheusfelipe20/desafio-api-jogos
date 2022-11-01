package funcoes

import "testing"

type cases struct {
	entrada  string
	esperado bool
}

type caseFormated struct {
	entrada  string
	esperado string
}

// teste para verificar se o a função que chama a verificação de cpf está funcionando
func Test_VerificarCPF(t *testing.T) {

	testesCases := []cases{
		{"", false},              // teste com cpf vazio
		{"699.602.50-16", false}, // teste com cpf com tamanho diferente de 11
		{"11111111111", false},   // teste com cpf com todos os números iguais
		{"76470909074", false},   // teste com cpf com o primeiro digito verificador inválido
		{"35129746009", false},   // teste com cpf com o segundo digito verificador inválido
		{"423.101.640-20", true}, // teste com cpf válido (com pontos e traço)
		{"79133406065", true},    // teste com cpf válido (sem pontos e traço)
	}

	for _, teste := range testesCases {
		resultado, _ := VerificarCPFbyString(teste.entrada)

		if resultado != teste.esperado {
			t.Errorf("Resultado %t diferente do esperado %t", resultado, teste.esperado)
		}
	}
}

// teste para verificar se a formatação do cpf está funcionando
func Test_Formatacao(t *testing.T) {

	testesCases := []caseFormated{
		{"836.047.617-94", "83604761794"}, // teste separando pontos e traço
		{"368067929-79", "36806792979"},   // teste separando traço
		{"461.103.34499", "46110334499"},  // teste separando pontos
	}

	for _, teste := range testesCases {
		resultado := Formated(teste.entrada)

		if resultado != teste.esperado {
			t.Errorf("Resultado %s diferente do esperado %s", resultado, teste.esperado)
		}
	}
}
