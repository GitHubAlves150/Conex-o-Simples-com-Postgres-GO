package handdlers

import (
	"gorm.io/gorm"
)

// Criamos a struct do controlador. O "db" privado é a nossa dependência.
type UserHandler struct {
	db *gorm.DB
}

// Esta função configura um novo UserHanddler (injeção de dependẽncia)
func NewUseHanddler(database *gorm.DB) *UserHandler {
	return &UserHandler{
		db: database,
	}

}
