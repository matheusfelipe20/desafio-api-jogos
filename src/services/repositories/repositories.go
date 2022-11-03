package repositories

import (
	"database/sql"

	"github.com/matheusfelipe20/projeto-api-jogos/src/Models"
)

type vendas struct {
	db *sql.DB
}

// NovoRepositorioDeVendas irá criar um repositório de vendas
func NovoRepositorioDeVendas(db *sql.DB) *vendas {
	return &vendas{db}
}

// CriarVenda irá inseir uma venda no banco de dados
func (v vendas) CriarVenda(venda models.Vendas) (int, error) {

	stmt, err := v.db.Prepare("insert into venda (id_jogo, titulo_jogo, campeonato, data_jogo, opcao_aposta, opcao_valor, valor_aposta, limite_aposta, cliente_nome, cliente_cpf, cliente_nascimento, ganho_provavel) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int
	if err = stmt.QueryRow(venda.Id_jogo, venda.Titulo_jogo, venda.Campeonato, venda.Data_jogo, venda.Opcao_aposta, venda.Opcao_valor, venda.Valor_aposta, venda.Limite_aposta, venda.Cliente_nome, venda.Cliente_cpf, venda.Cliente_nascimento, venda.Ganho_provavel).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// BuscarVendas irá retornar todas as vendas cadastradas no banco de dados
func (v vendas) BuscarVendas() ([]models.Vendas, error) {
	rows, err := v.db.Query("select * from venda")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []models.Vendas
	for rows.Next() {
		var venda models.Vendas

		if err = rows.Scan(&venda.Id, &venda.Id_jogo, &venda.Titulo_jogo, &venda.Campeonato, &venda.Data_jogo, &venda.Opcao_aposta, &venda.Opcao_valor, &venda.Valor_aposta, &venda.Limite_aposta, &venda.Cliente_nome, &venda.Cliente_cpf, &venda.Cliente_nascimento, &venda.Ganho_provavel); err != nil {
			return nil, err
		}
		vendas = append(vendas, venda)
	}

	return vendas, nil
}