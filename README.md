# 📖 Documentação do Projeto: API em Go (Clean Architecture & SOLID)

Este documento explica o funcionamento de cada arquivo do projeto e como eles cooperam para realizar a busca de dados no PostgreSQL (método GET) utilizando o **Chi** e o **GORM**.

---

## 🏎️ A Analogia do Carro e do Motor

Para entender a arquitetura deste projeto, usamos uma analogia simples:
* **`UserHandler` (O Carro):** É a estrutura responsável por se movimentar (receber pedidos da web e entregar respostas).
* **`gorm.DB` (O Motor):** É o que dá força ao carro para que ele funcione. Sem o motor, o carro não consegue ir até à base de dados buscar as informações.
* **Injeção de Dependência:** O carro não fabrica o seu próprio motor. O motor é fabricado na oficina (`main.go`) e é **injetado (encaixado)** dentro do carro.

---

## 📂 Explicação Arquivo por Arquivo

### 1. `main.go` (A Oficina / O Maestro)
Este é o ponto de partida de toda a aplicação. Ele funciona como a oficina que liga os componentes e dá a partida no servidor.

```go
package main

import (
	"api_teste/internal/handdlers"
	"api_teste/internal/repository"
	"fmt"
	"net/http"

	"://github.com"
)

func main() {
	// 1. Cria o Motor: Liga-se ao banco de dados e traz a conexão pronta.
	db := repository.ConectDB()
	
	// 2. Monta o Carro: Cria o controlador inserindo o motor (db) dentro dele.
	userHandler := handdlers.NewUseHanddler(db)

	// 3. Define as Estradas: Configura o roteador Chi.
	r := chi.NewRouter()

	// 4. Cria a Rota: Diz que quando alguém aceder a "/users" via GET, o controlador age.
	r.Get("/users", userHandler.ObterUsuarios)
	
	// 5. Liga o Carro: Inicia o servidor HTTP na porta 8080.
	fmt.Println("Servidor rodando na porta :8080 em http://localhost")
	http.ListenAndServe(":8080", r)
}
```

---

### 2. `internal/repository/db.go` (A Fábrica de Motores)
Este arquivo tem apenas **uma responsabilidade** (S do SOLID): sabe como falar com o PostgreSQL e criar o motor (`db`).

```go
package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectDB() *gorm.DB {
	dsn := "host=localhost user=usuario_lucas password=123456 dbname=banco_teste_lucas port=5432 sslmode=disable"

	// Abre a ligação técnica com o Postgres usando o GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao ligar ao banco de dados")
	}
	return db // Devolve o motor pronto
}
```

---

### 3. `internal/handdlers/use_handler.go` (O Design do Carro e Encaixe do Motor)
Aqui nós definimos o espaço para guardar o motor dentro do carro e criamos a função de montagem.

```go
package handdlers

import (
	"gorm.io/gorm"
)

// UserHandler representa o nosso Carro.
type UserHandler struct {
	db *gorm.DB // Este campo privado é o espaço reservado para o Motor.
}

// NewUseHanddler é a linha de montagem do carro (Injeção de Dependência).
// Recebe um banco de fora (database) e encaixa-o no espaço "db" do carro.
func NewUseHanddler(database *gorm.DB) *UserHandler {
	return &UserHandler{
		db: database, // O encaixe do motor ocorre aqui.
	}
}
```

---

### 4. `internal/handdlers/method_GET.go` (O Carro em Movimento / A Ação)
Este arquivo é o que executa a ação toda vez que um cliente faz uma requisição HTTP GET para a rota `/users`.

```go
package handdlers

import (
	"api_teste/internal/estrutura"
	"encoding/json"
	"net/http"
)

// ObterUsuarios usa o motor do carro (h.db) para acelerar e trazer os dados.
func (h *UserHandler) ObterUsuarios(w http.ResponseWriter, r *http.Request) {
	// Passo 1: Cria um balde vazio na memória que espera uma lista de estruturas.
	var users []estrutura.UserResponse

	// Passo 2: O GORM usa o motor (h.db) para fazer um "SELECT * FROM users" 
	// e enche o nosso balde (&users) com os dados reais vindos do Postgres.
	err := h.db.Find(&users).Error
	if err != nil {
		http.Error(w, "Erro ao conectar-se com o banco", http.StatusInternalServerError)
		return
	}

	// Passo 3: Avisa o cliente que a resposta deu certo (200 OK) e será um JSON.
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Passo 4: Transforma o balde cheio de structs em formato JSON (texto)
	// e envia diretamente para a tela do utilizador através do escritor (w).
	json.NewEncoder(w).Encode(users)
}
```

---

### 5. `internal/estrutura/user.go` (O Molde dos Dados)
Este arquivo define o formato exato de como os dados saem do banco e como se transformam em texto JSON. Ele mapeia as colunas do Postgres para o Go.

```go
package estrutura

import (
	"time"
)

type UserResponse struct {
	Id              string    `json:"id" gorm:"primaryKey"`
	Nome            string    `json:"nome"`
	Email           string    `json:"email"`
	Senha           string    `json:"-"` // O "-" esconde a senha, impedindo que ela seja enviada no GET
	Medicamento     string    `json:"medicamento"`
	Data_nascimento time.Time `json:"data_nascimento" gorm:"type:date"` // Mapeia o DATE do Postgres
	Criado          time.Time `json:"criadoEm" gorm:"column:criado"`    // Mapeia a coluna "criado"
}

// TableName força o GORM a apontar exatamente para a tabela "users" em minúsculo.
func (UserResponse) TableName() string {
	return "users"
}
```

---

## 🔄 Resumo do Fluxo na Execução (Linha do Tempo)

1. Você executa o servidor (`main.go`).
2. O **`repository`** cria o motor da base de dados.
3. O **`use_handler.go`** cria o controlador injetando o motor nele.
4. O cliente faz um pedido HTTP GET para `/users`.
5. O Chi direciona o pedido para o arquivo **`method_GET.go`**.
6. O código cria um array baseado no molde do arquivo **`estrutura`**.
7. O GORM consulta a base de dados, preenche o array e envia a resposta convertida em JSON para o cliente.


___

Este conteúdo é parte do meu aprendizado em Golang. Nesta etapa estou unindo as tecnologia que já havia aprendido em outras área como HTML,CSS, JS, Banco de Dado. Agora estou documentando pequenos fragmentos que compoêm um sistema web como neste caso trabalhar com API utilizando as rotas principais /GET, /PUT, /DELETE, /POST trabalhando diretamente com o banco. Todo este aprendizado está sendo possível com ajuda de I.A que me orienta com as melhores técnicas de programação.