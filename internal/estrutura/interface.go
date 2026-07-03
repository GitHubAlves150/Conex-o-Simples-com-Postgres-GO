package estrutura

import (
	"time"
)

// UserResponse representa o formato de saída dos dados
type UserResponse struct {
	Id              string    `json:"id" gorm:"primaryKey"`
	Nome            string    `json:"nome"`
	Email           string    `json:"email"`
	Senha           string    `json:"-"`
	Medicamento     string    `json:"medicamento"`
	Data_nascimento time.Time `json:"data_nascimento" gorm:"type:date"`
	Criado          time.Time `json:"criadoEm" gorm:"column:criado"`
}

func (UserResponse) TableName() string {
	return "users"
}


