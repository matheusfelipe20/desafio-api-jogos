package database

import (
	"database/sql"

	_ "github.com/lib/pq" // driver de conexão do postgres
)

// Connect abre a conexão com o postgres
func Connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", stringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
