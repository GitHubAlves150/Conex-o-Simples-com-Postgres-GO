package main

import (
	"app/internal/handdler"
	"app/internal/interface_test"
	"app/internal/repository"
	"app/internal/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors" // <--- 1. ADICIONADO O IMPORT DO CORS AQUI!
)

func main() {
	db, err := repository.OpenDBWithGORM()
	if err != nil {
		log.Fatalf("Falha crítica no banco: %v", err)
	}

	err = db.AutoMigrate(&interface_test.User{})
	if err != nil {
		log.Fatalf("Erro ao rodar migrações: %v", err)
	}

	repo := repository.NewGormRepository(db)
	srv := service.NewUsuarioService(repo)
	handler := handdler.NewUsuarioHandler(srv)

	router := chi.NewRouter()
	// 1. IMPORTANTE: Adicione as configurações de CORS antes de definir a rota!
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permite requisições de qualquer origem (ideal para teste local)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache de pre-flight em segundos
	}))

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Post("/usuary", handler.CriaUsuarioHanddler)

	fmt.Println("Servidor iniciado em http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
