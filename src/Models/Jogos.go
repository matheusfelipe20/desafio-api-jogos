package models

type Jogo struct {
	ID            int                  `json:"id,omitempty"`
	Titulo        string               `json:"titulo,omitempty"`
	ID_Campeonato int                  `json:"id_campeonato,omitempty"`
	Data          string               `json:"data_jogo,omitempty"`
	Opcoes        []map[string]float64 `json:"opcoes,omitempty"`
	Limite        []map[string]float64 `json:"limite,omitempty"`
}
