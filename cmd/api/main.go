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
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Post("/usuary", handler.CriaUsuarioHanddler)

	fmt.Println("Servidor iniciado em http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}