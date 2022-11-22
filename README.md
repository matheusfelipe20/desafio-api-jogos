<h1 align="center">Desafio Api Jogos🎲</h1>

## 💻 Descrição do projeto

O projeto tem como objetivo simular um sistema de apostas esportivas, ele deve consumir de uma API contendo os campeonatos, jogos e CPFs dos usuários já cadastrados. Além disso, também deve ser possível consultar na sua API todos os eventos ou filtrar.

---

## 🔨 Funcionalidades:

- Realizar apostas: POST `/vendas`;
- Consultar todas as apostas realizadas: GET `/vendas`;
- Consultar os jogos disponíveis: GET `/eventos`;
- Consultar os campeonatos disponíveis: GET `/campeonatos`;
- Consultar informações dos usuários do sistema: GET `/cpf/{cpf}`;
- Filtrar eventos por: GET `/eventos/{id}`, `/eventos/campeonato/{id}` e `/eventos/data/{data}`.

---

## :receipt: Estrutura para realizar o POST '/vendas':

```
{
  "id_jogo": 354858757161272,
  "opcao_aposta": "1",
  "valor_aposta": 100,
  "cliente_cpf": "368.067.929-79"
}
```
- No primeiro campo será inserido o id do jogo escolhido;
- Posteriormente será necessário informar qual a opção da aposta, nesse caso há três opções: ( 1 : casa | x : empate | 2 : fora);
- No "valor_aposta" será preciso colocar o valor da aposta do usuário, porém cuidado com o limite, pois cada opção e jogo há um limite para o valor de cada aposta;
- E por último o cpf precisa ser válido e está cadastrado no sistema;
