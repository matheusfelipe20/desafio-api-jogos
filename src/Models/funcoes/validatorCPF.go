package funcoes

import (
	"strconv"
	"strings"
)

// VerificaCPF irá verificar se o cpf é válido
func VerificarCPF(cpf int) bool {
	cpfString := strconv.Itoa(cpf)

	if !Tamanho(cpfString) || VerificarNumerosIguais(cpfString) {
		return false
	} else if VerificacaoPorDigito(Numbers(cpfString)) {
		return true
	} else {
		return false
	}

}

// função para verificar se um cpf do tipo string é válido
func VerificarCPFbyString(cpf string) (bool, int) {

	var cpfGerado int
	cpfGerado, _ = strconv.Atoi(Formated(cpf))

	if !Tamanho(cpf) || VerificarNumerosIguais(cpf) {
		return false, 0
	} else if VerificacaoPorDigito(Numbers(cpf)) {
		return true, cpfGerado
	} else {
		return false, 0
	}

}

// função para verificar o tamanho de um cpf
func Tamanho(cpf string) bool {
	if len(Formated(cpf)) != 11 {
		return false
	} else {
		return true
	}
}

// função que retorna o cpf sem os caracteres especiais
func Formated(cpf string) string {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	return cpf
}

// função para converter os caracteres de uma string cpf para um slice de int
func Numbers(cpf string) []int {
	separados := []int{}
	cpfFormated := Formated(cpf)

	var digitoInNumber int64
	for _, letra := range cpfFormated {
		digitoInNumber, _ = strconv.ParseInt(string(letra), 10, 64)
		separados = append(separados, int(digitoInNumber))
	}

	return separados
}

// função para verificar se os caracteres de um cpf são iguais
func VerificarNumerosIguais(cpf string) bool {

	cpfFormated := Formated(cpf)

	if cpfFormated == "00000000000" || cpfFormated == "11111111111" ||
		cpfFormated == "22222222222" || cpfFormated == "33333333333" ||
		cpfFormated == "44444444444" || cpfFormated == "55555555555" ||
		cpfFormated == "66666666666" || cpfFormated == "77777777777" ||
		cpfFormated == "88888888888" || cpfFormated == "99999999999" {
		return true
	}

	return false
}

// função para fazer a verificação por digitoa verificadores do cpf
func VerificacaoPorDigito(cpfS []int) bool {

	digito1 := cpfS[0]
	digito2 := cpfS[1]
	digito3 := cpfS[2]
	digito4 := cpfS[3]
	digito5 := cpfS[4]
	digito6 := cpfS[5]
	digito7 := cpfS[6]
	digito8 := cpfS[7]
	digito9 := cpfS[8]
	primeiroDigito := cpfS[9]
	segundoDigito := cpfS[10]

	algoritmoPrimeiroDigito := (digito1*10 + digito2*9 + digito3*8 + digito4*7 + digito5*6 + digito6*5 + digito7*4 + digito8*3 + digito9*2) * 10 % 11
	if algoritmoPrimeiroDigito == 10 {
		algoritmoPrimeiroDigito = 0
	}

	algoritmoSegundoDigito := (digito1*11 + digito2*10 + digito3*9 + digito4*8 + digito5*7 + digito6*6 + digito7*5 + digito8*4 + digito9*3 + primeiroDigito*2) * 10 % 11
	if algoritmoSegundoDigito == 10 {
		algoritmoSegundoDigito = 0
	}

	if (algoritmoPrimeiroDigito == primeiroDigito) && (algoritmoSegundoDigito == segundoDigito) {
		return true
	} else {
		return false
	}
}
