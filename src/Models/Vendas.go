package models

import (
	"errors"

	"github.com/matheusfelipe20/projeto-api-jogos/src/Models/funcoes"
)

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
		return errors.New("id do jogo é igual a 0")
	}
	if verificarTituloJogo := funcoes.ValidarCampo(vd.Titulo_jogo); !verificarTituloJogo {
		return errors.New("falha ao  cadastrar, insira o titulo do jogo")
	}
	if verificarCampeonato := funcoes.ValidarCampo(vd.Campeonato); !verificarCampeonato {
		return errors.New("falha ao cadastrar, insira o campeonato")
	}
	if verificarDataJogo := funcoes.ValidadeDataVenda(vd.Data_jogo); verificarDataJogo {
		return errors.New("falha ao cadastrar, insira a data do jogo, ou verfique se o jogo ainda está disponivel")
	}
	if verificarNomeCliente := funcoes.ValidarCampo(vd.Cliente_nome); !verificarNomeCliente {
		return errors.New("falha ao cadastrar, insira o nome do cliente")
	}
	if verificarCpfCliente, _ := funcoes.VerificarCPFbyString(vd.Cliente_cpf); !verificarCpfCliente {
		return errors.New("falha ao cadastrar, cpf inválido")
	}
	if verificarNomeCliente := funcoes.ValidadeDataNascimento(vd.Cliente_nascimento); !verificarNomeCliente {
		return errors.New("falha ao cadastrar, usuário menor de idade")
	}
	if verificarLimiteAposta := funcoes.ValidadeLimiteValor(vd.Limite_aposta, vd.Valor_aposta); !verificarLimiteAposta {
		return errors.New("falha ao cadastrar, valor da aposta insuficiente, ou excedeu o limite do valor da aposta")
	}

	return nil
}
