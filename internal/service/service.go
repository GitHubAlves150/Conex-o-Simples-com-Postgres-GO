package service

import (
	"app/internal/interface_test"
	"app/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UsuarioService interface {
	CriarUsuario(nome, email, senha string) (*interface_test.UserResponse, error)
}

type Servico struct {
	repo repository.UsuarioRepository //interface de salvar o banco
}

func NewUsuarioService(repo repository.UsuarioRepository) *Servico {
	return &Servico{repo: repo}
}

func (s *Servico) CriarUsuario(nome, email, senha string) (*interface_test.UserResponse, error) {
	if nome == "" {
		return nil, errors.New("nome não pode ser vazio")
	}

	novoUsuario := &interface_test.User{
		ID:       uuid.New().String(),
		Name:     nome,
		Email:    email,
		Senha:    senha,
		CriadoEm: time.Now(),
	}

	err := s.repo.SalvarNoBanco(novoUsuario)
	if err != nil {
		return nil, err
	}

	return &interface_test.UserResponse{
		ID:        novoUsuario.ID,
		Name:      novoUsuario.Name,
		Email:     novoUsuario.Email,
		Senha:     novoUsuario.Senha,
		Criado_em: novoUsuario.CriadoEm,
	}, nil
}