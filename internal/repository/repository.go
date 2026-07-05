package repository

import (
	"app/internal/interface_test"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	SalvarNoBanco(user *interface_test.User) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) SalvarNoBanco(user *interface_test.User) error {
	result := r.db.Create(user)
	return result.Error
}