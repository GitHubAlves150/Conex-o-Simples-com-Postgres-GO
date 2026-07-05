package repository

import (
	"app/internal/utility"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBWithGORM() (*gorm.DB, error) {
	dsn := utility.CredenciaisDB()
	// O GORM abre a conexão usando o driver interno dele
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com GORM: %v", err)
	}

	// Para configurar o Pool de Conexões, pegamos o *sql.DB nativo por baixo do GORM
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configurações do Pool obrigatórias para o mercado
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Pool de conexões via GORM inicializado!")
	return db, nil
}
