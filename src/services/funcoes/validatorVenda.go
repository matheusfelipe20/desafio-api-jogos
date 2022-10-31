package funcoes

import (
	"log"
	"time"
)

//Validador para limitar a venda do jogo a 1 dia antes do evento
func ValidadeDataVenda(data string) bool {

	if data == "" {
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

func ValidadeLimiteValor(limite_aposta float64, valor_aposta float64) bool {

	return limite_aposta >= valor_aposta

}
