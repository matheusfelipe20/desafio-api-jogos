package models

import (
	"fmt"

	"github.com/matheusfelipe20/projeto-api-jogos/src/Models/funcoes"
)

// struct para teste de vendas
type RespVenda struct {
	Code    string
	Message string
}

type ErroVenda struct {
	Code string
	Err  string
}

func (r *ErroVenda) Error() string {
	return fmt.Sprintf("StatusCode: %v, Erro: %v", r.Code, r.Err)
}

type Vendas struct {
	Id                 int     `json:"id,omitempty"`
	Id_jogo            int     `json:"id_jogo,omitempty"`
	Titulo_jogo        string  `json:"titulo_jogo,omitempty"`
	Campeonato         string  `json:"campeonato,omitempty"`
	Data_jogo          string  `json:"data_jogo,omitempty"`
	Opcao_aposta       string  `json:"opcao_aposta,omitempty"`
	Opcao_valor        float64 `json:"opcao_valor,omitempty"`
	Valor_aposta       float64 `json:"valor_aposta,omitempty"`
	Limite_aposta      float64 `json:"limite_aposta,omitempty"`
	Cliente_nome       string  `json:"cliente_nome,omitempty"`
	Cliente_cpf        string  `json:"cliente_cpf,omitempty"`
	Cliente_nascimento string  `json:"cliente_nascimento,omitempty"`
	Ganho_provavel     float64 `json:"ganho_provavel,omitempty"`
}

func (vd *Vendas) ValidarVendas() error {

	if verificarIdJogo := funcoes.ValidadeID(uint64(vd.Id_jogo)); !verificarIdJogo {
		message := "ID de jogo não encontrado"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}

	if verificarTituloJogo := funcoes.ValidarCampo(vd.Titulo_jogo); !verificarTituloJogo {
		message := "titulo de jogo inválido"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}

	if verificarCampeonato := funcoes.ValidarCampo(vd.Campeonato); !verificarCampeonato {
		message := "campeonato inválido"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}

	if verficarDataJogo := funcoes.ValidadeDataVenda(vd.Data_jogo); verficarDataJogo {
		message := "jogo indisponível, horário para apostar ultrapassado"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}
	if verificarCpfCliente, _ := funcoes.VerificarCPFbyString(vd.Cliente_cpf); !verificarCpfCliente {
		message := "CPF inválido"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}
	if verificarNomeCliente := funcoes.ValidarCampo(vd.Cliente_nome); !verificarNomeCliente {
		message := "nome do cliente não encontrado"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}
	if verificarNomeCliente := funcoes.ValidadeDataNascimento(vd.Cliente_nascimento); !verificarNomeCliente {
		message := "usuário menor de idade"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}
	if verificarLimiteAposta := funcoes.ValidadeLimiteValor(vd.Limite_aposta, vd.Valor_aposta); !verificarLimiteAposta {
		message := "Limite do valor da aposta excedido"
		return &ErroVenda{
			Code: "400",
			Err:  message,
		}
	}

	return nil
}
