package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectDB() *gorm.DB {
	dsn := "host=localhost user=usuario_lucas password=123456 dbname=banco_teste_lucas port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao ligr ao banco de dados")
	}
	return db
}
