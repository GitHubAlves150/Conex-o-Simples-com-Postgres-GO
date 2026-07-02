package repository

import (
	"app/internal/estrutura"
	"fmt"
	"time"
)

func SalvarAT_DB(id, nome, email, senha string, criado time.Time) (*estrutura.UserResponse, error) {

	db, err := OpenDB()
	if err != nil {
		return db, nil
	}

	defer db.Close()
	query := "INSERT INTO users (id, nome, email, senha, criado) VALUES ($1, $2, $3, $4, $5)"
	_, ER := db.Exec(query, id, nome, email, senha, criado)
	if ER != nil {
		return nil, fmt.Errorf("Erro ao salvar usuário no banco de dados: %v", err)
	}

	user := &estrutura.UserResponse{
		ID:        id,
		Name:      nome,
		Email:     email,
		Senha:     senha,
		Criado_em: time.Now(),
	}

	return user, nil

}

go get github.com/google/uuid
go get github.com/lib/pq
go get github.com/go-chi/chi/v5

curl -X POST http://localhost:8080/Perfil \
  -H "Content-Type: application/json" \
  -d '{"nome": "Lucas", "email": "lucas@email.com", "senha": "123", "medicamento":"Albatroz de cavilatacio", "data_nascimento": "1988-08-31"}'

