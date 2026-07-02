package main

import (
	"app/internal/handdler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Post("/usuary", handdler.CriaUsuarioHanddler)

	fmt.Println("Servidor iniciado em http://localhost:8080...")
	http.ListenAndServe(":8080", router) // Liga o servidor à porta 8080 usando as rotas do Chi

}
