package funcoes

// ValidadeLimiteValor verifica se o valor da aposta é menor ou igual ao limite
func ValidadeLimiteValor(limite_aposta float64, valor_aposta float64) bool {

	return limite_aposta >= valor_aposta

}
