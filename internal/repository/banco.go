package repository

import (
	"app/internal/interface_test"
	"fmt"
	"time"
)

func SalvarAT_DB(id, nome, email, senha string, criado time.Time) (*interface_test.UserResponse, error) {

	db, err := OpenDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	query := "INSERT INTO users (id, nome, email, senha, criado) VALUES ($1, $2, $3, $4, $5)"
	_, ER := db.Exec(query, id, nome, email, senha, criado)
	if ER != nil {
		return nil, fmt.Errorf("Erro ao salvar usuário no banco de dados: %v", ER)
	}

	user := &interface_test.UserResponse{
		ID:        id,
		Name:      nome,
		Email:     email,
		Senha:     senha,
		Criado_em: time.Now(),
	}

	return user, nil

}
