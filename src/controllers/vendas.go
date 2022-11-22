package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	models "github.com/matheusfelipe20/projeto-api-jogos/src/Models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/database"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/api"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/odds"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/repositories"
	"github.com/matheusfelipe20/projeto-api-jogos/src/services/respostas"
)

// RealizarVenda irá realizar uma venda e cadastrar no banco de dados
func RealizarVenda(w http.ResponseWriter, r *http.Request) {

	// lendo o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Criando um objeto do tipo venda
	var venda models.Vendas
	err = json.Unmarshal(requestBody, &venda)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	// adicionando as opeações da aposta
	for _, jogo := range api.GetJogos() {
		// adicionando as odds nas informações de venda
		if jogo.ID == venda.Id_jogo {
			_, value := odds.CompararOpcaoAposta(venda.Opcao_aposta, jogo.Opcoes) // opção de odd do jogo
			venda.Opcao_valor = value                                             // atribuindo o valor da opção de aposta à venda
			val := odds.CompararLimiteComOpcao(venda.Opcao_aposta, jogo.Limites)  // limite de odd do jogo
			venda.Limite_aposta = val                                             // atribuindo o limite de aposta à venda

			ganho := odds.CalcularGanho(venda.Valor_aposta, venda.Opcao_valor)
			venda.Ganho_provavel = ganho // atribuindo o ganho provável à venda

			venda.Titulo_jogo = jogo.Titulo // atribuindo o título do jogo à venda
			venda.Data_jogo = jogo.Data     // atribuindo a data do jogo à venda

			for _, campeonato := range api.GetCampeonatos() {
				if campeonato.ID == jogo.ID_Campeonato {
					venda.Campeonato = campeonato.Titulo // atribuindo o campeonato do jogo à venda
				}
			}
			break
		}
	}

	for _, usuario := range api.GetUser() {
		if usuario.Cpf == venda.Cliente_cpf {
			venda.Cliente_nome = usuario.Nome             // atribuindo o nome do cliente à venda
			venda.Cliente_nascimento = usuario.Nascimento // atribuindo a data de nascimento do cliente à venda
			break
		}
	}

	// fazendo validação de venda
	val := venda.ValidarVendas()
	if val != nil {
		respostas.Erro(w, http.StatusBadRequest, val)
		return
	}

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// registrando o objeto da venda no repositório
	repositorio := repositories.NovoRepositorioDeVendas(db)
	venda.Id, err = repositorio.CriarVenda(venda)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resultadoOK := "Aposta realizada com sucesso!"
	respVenda := models.RespVenda{Code: "200", Message: resultadoOK}
	respostas.JSON(w, http.StatusCreated, respVenda)
}

// ListarVendas irá buscar todas as vendas feitas
func ListarVendas(w http.ResponseWriter, r *http.Request) {

	// abrindo conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// buscando todas as vendas no repositório
	repositorio := repositories.NovoRepositorioDeVendas(db)
	vendas, err := repositorio.BuscarVendas()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// retornando as vendas para o usuário
	respostas.JSON(w, http.StatusOK, vendas)

}
