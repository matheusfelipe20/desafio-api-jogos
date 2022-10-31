package funcoes

import (
	"log"
	"time"
)

//Validador para verificar se o usuario é maior de idade
func ValidadeDataNascimento(nascimento string) bool {

	if nascimento == "" {
		return false
	}

	parsed, err := time.Parse("02/01/2006", nascimento)
	if err != nil {
		log.Println(err)
	}
	beforeYear := parsed.AddDate(18, 0, 0) //Somar 18 anos a data de nascimento do usuario
	today := time.Now()

	compareted := beforeYear.Before(today)

	return compareted
}
