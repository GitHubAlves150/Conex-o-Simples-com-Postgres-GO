package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {

	dsn := "host=localhost port=5432 user=usuario_lucas password=123456 dbname=banco_teste_lucas sslmode=disable"

	//inicializa o pool de conexoes
	db, err := sql.Open("postgres", dsn)

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco realizada com sucesso")
	return db, nil
}
