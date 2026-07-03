package main

import (
	"api_teste/internal/handdlers"
	"api_teste/internal/repository"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	db := repository.ConectDB()
	userHandler := handdlers.NewUseHanddler(db)

	//configura o roteador chi
	r := chi.NewRouter()

	//registra a rota. Passamos o método da struct que criamos acima
	r.Get("/users", userHandler.ObterUsuarios)
	
	//Inicia o servidor
	fmt.Println("Servidor rodando na porta :8080 em hhtp//localhost")
	http.ListenAndServe(":8080", r)

}

