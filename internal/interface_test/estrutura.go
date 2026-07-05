package interface_test

import "time"

// Struct que representa a tabela no banco de dados para o GORM
type User struct {
	ID       string    `gorm:"primaryKey;type:uuid"`
	Name     string    `gorm:"column:nome;type:varchar(100);not null"`
	Email    string    `gorm:"uniqueIndex;type:varchar(100);not null"`
	Senha    string    `gorm:"type:varchar(255);not null"`
	CriadoEm time.Time `gorm:"column:criado;default:CURRENT_TIMESTAMP"`
}

// DTOs (Data Transfer Objects) para a API HTTP continuar igual
type UserRequest struct {
	Name  string `json:"name"`
	Senha string `json:"senha"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Senha     string    `json:"senha"`
	Criado_em time.Time `json:"criado_em"`
}