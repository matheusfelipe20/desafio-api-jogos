package funcoes

// Verifica se o ID é diferente de zero
func ValidadeID(id uint64) bool {
	idJogo := id

	return idJogo != 0
}

// Verifica se o campo do título/campeonato do jogo está vazio
func ValidarCampo(titulo string) bool {
	return titulo != ""
}
