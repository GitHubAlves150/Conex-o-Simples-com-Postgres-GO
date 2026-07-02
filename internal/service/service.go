package service

import (
	"app/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CriarUsuarioService(nome, email, senha string) (*UserResponse, error) {
	if nome == "" {
		return nil, errors.New("Nome não pode ser vazio")
	}

	id := uuid.New().String()
	fmt.Println("ID gerado:", id)
	//salva no banco

	user, err := repository.SalvarAT_DB(id, nome, email, senha, time.Now())
	if err != nil {
		return nil, err
	}	

	return user, nil

}
