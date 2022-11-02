package funcoes

import (
	"log"
	"time"
)

//Validador para limitar a venda do jogo a 1 dia antes do evento
func ValidadeDataVenda(data string) bool {

	if data == "" || data == "0000-00-00" {
		return false
	}
	if data < "2022-10-31" || data > "2022-11-01" {
		return false
	}

	parsed, err := time.Parse("2006-01-02", data)
	if err != nil {
		log.Println(err)
	}
	beforeDay := parsed.AddDate(0, 0, 1) //Somar 1 dia a data do evento
	today := time.Now()

	comparetedEvento := beforeDay.Before(today)

	return comparetedEvento
}
