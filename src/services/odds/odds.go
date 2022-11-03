package odds

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
