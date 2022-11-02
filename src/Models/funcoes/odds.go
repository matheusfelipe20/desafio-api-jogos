package funcoes

func CalcularGanho(aposta float64, valorOpcao float64) float64 {

	return aposta * valorOpcao
}

func CompararOpcaoAposta(opcaoAposta string, opcoes []map[string]float64) (string, float64) {

	for _, opcao := range opcoes {
		for key, value := range opcao {
			if key == opcaoAposta {
				return key, value
			}
		}
	}
	return "", 0.0
}

// CompararLimiteComOpcao compara o limite com a opção selecionada pelo usuário e retorna o limite referente a opção
func CompararLimiteComOpcao(opcaoAposta string, limite []map[string]float64) float64 {

	for _, opcao := range limite {
		for key, value := range opcao {
			if key == opcaoAposta {
				return value
			}
		}
	}
	return 0.0
}

// funções de implementações das odds e limites
func MapOpcoes(a float64, b float64, c float64) []map[string]float64 {
	opcoes := []map[string]float64{}

	opcoes = append(opcoes, map[string]float64{"casa": a})
	opcoes = append(opcoes, map[string]float64{"empate": b})
	opcoes = append(opcoes, map[string]float64{"fora": c})

	return opcoes
}

func MapLimites(a float64, b float64, c float64) []map[string]float64 {
	limite := []map[string]float64{}

	limite = append(limite, map[string]float64{"casa": a})
	limite = append(limite, map[string]float64{"empate": b})
	limite = append(limite, map[string]float64{"fora": c})

	return limite
}

func ValidadeLimiteValor(limite_aposta float64, valor_aposta float64) bool {

	return limite_aposta >= valor_aposta

}
