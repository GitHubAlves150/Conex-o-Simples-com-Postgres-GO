# 🚀 API Go com PostgreSQL & Docker

Uma API REST simples e estruturada desenvolvida em Go (Golang) para fins de estudo. O projeto demonstra a criação de uma arquitetura em camadas (Handler, Service, Repository) conectada a uma base de dados PostgreSQL isolada num contentor Docker.

## 🛠️ Tecnologias Utilizadas

* **Go (Golang)** (v1.23+)
* **PostgreSQL** (v15)
* **Docker & Docker Compose** (Para orquestração do banco de dados)
* **Go-Chi** (Router HTTP leve)
* **Google UUID** (Para geração de IDs únicos)
* **Lib/PQ** (Driver nativo do Postgres para Go)

## 📁 Estrutura do Projeto

```text
APITeste/
├── cmd/
│   └── api/
│       └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── estrutura/
│   │   └── user.go           # Structs de Request e Response (Modelos)
│   ├── handdlerUser/
│   │   └── user_handler.go   # Camada de controle HTTP (JSON/Status Code)
│   ├── repository/
│   │   ├── conectDB.go       # Inicialização e Ping do banco de dados
│   │   └── salvarUser_DB.go  # Queries SQL (Inserts/Exec)
│   └── servico/
│       └── criar_user.go     # Regras de negócio e geração de UUID
├── docker-compose.yml        # Configuração do container do Postgres
└── go.mod                    # Gestão de dependências do Go
```

## 🚀 Como Executar o Projeto

### 1. Clonar o repositório
```bash
git clone https://github.com
cd NOME_DO_REPOSITORIO
```

### 2. Iniciar o Banco de Dados (Docker)
Certifique-se de que tem o Docker instalado e execute:
```bash
docker compose up -d
```

### 3. Configurar a Tabela no Postgres
Entre no terminal do Postgres dentro do container:
```bash
docker exec -it postgres_teste psql -U usuario_lucas -d banco_teste_lucas
```
Cole o seguinte script SQL para criar a tabela:
```sql
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    senha VARCHAR(100) NOT NULL,
    medicamento VARCHAR(30) NOT NULL,
    data_nascimento DATE,
    criado TIMESTAMPTZ DEFAULT NOW()
);
```
Digite `\q` para sair do terminal do Postgres.

### 4. Iniciar a API em Go
Na raiz do projeto, instale as dependências e rode o servidor:
```bash
go mod tidy
go run cmd/api/main.go
```
O servidor será iniciado em: `http://localhost:8080`

## 🧪 Como Testar (Endpoints)

### Criar Perfil de Utilizador
* **Rota:** `/Perfil`
* **Método:** `POST`
* **Headers:** `Content-Type: application/json`

**Exemplo de Requisição (cURL):**
```bash
curl -X POST http://localhost:8080/Perfil \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Lucas", 
    "email": "lucas@email.com", 
    "senha": "123", 
    "medicamento": "Albatroz de cavilatacio", 
    "data_nascimento": "1988-08-31"
  }'
```

---
# 📖 Documentação do Projeto: Primeira Etapa - Inserção de Dados (POST)

Este documento detalha o funcionamento da primeira etapa da API em Go, focada na criação e inserção de dados (método POST) no PostgreSQL utilizando o **Chi**, o driver nativo `database/sql` e o **UUID**.

---

## 🏎️ A Analogia do Carro e do Motor (Versão Fábrica)

Para entender como os dados entram no sistema através do método POST, adaptamos a nossa analogia para uma linha de montagem:
* **`CriaUsuarioHanddler` (A Portaria da Fábrica):** Recebe a matéria-prima (os dados em formato JSON que o cliente enviou), confere se o formato está correto e autoriza a entrada.
* **`CriarUsuarioService` (A Linha de Produção / Regras):** É o cérebro que decide o que fazer. Ele valida se o nome está preenchido, gera o número de série único (o ID do usuário usando UUID) e monta o objeto final.
* **`SalvarAT_DB` (A Prensa / Armazenamento):** É a máquina que pega no objeto montado e grava-o de forma permanente dentro do depósito (o banco de dados PostgreSQL).
* **`OpenDB` (O Gerador de Energia):** Abre e fecha a conexão com a energia (o banco de dados) toda vez que uma peça precisa de ser gravada.

---

## 📦 Explicação Arquivo por Arquivo

### 1. `main.go` (A Oficina / O Maestro)
Inicia o roteador e abre as portas da aplicação para ouvir requisições na rede.

```go
package main

import (
	"app/internal/handdler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// 1. Configura as estradas digitais da aplicação usando o Chi
	router := chi.NewRouter()

	// 2. Define que quando chegar um pedido POST em "/usuary", o guardião do handler entra em ação
	router.Post("/usuary", handdler.CriaUsuarioHanddler)

	// 3. Liga o servidor HTTP na porta 8080
	fmt.Println("Servidor iniciado em http://localhost:8080...")
	http.ListenAndServe(":8080", router) 
}
```

---

### 2. `internal/handdler/user_handler.go` (A Portaria / Validador HTTP)
Este arquivo recebe o pedido da web, lê o JSON que veio no corpo da requisição (`r.Body`) e repassa os dados limpos para a camada de negócios (Service).

```go
package handdler

import (
	"app/internal/service"
	"encoding/json"
	"net/http"
)

func CriaUsuarioHanddler(w http.ResponseWriter, r *http.Request) {
	var Req estrutura.UserResponse // Cria o molde vazio para receber os dados

	// Converte o JSON que o cliente enviou em uma struct que o Go entende
	err := json.NewDecoder(r.Body).Decode(&Req)
	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Envia os dados extraídos para o Service validar e processar
	user, err := service.CriarUsuarioService(Req.Nome, Req.Senha, Req.Email)
	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Devolve os dados do usuário recém-criado em formato JSON com o código 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
```

---

### 3. `internal/service/user_service.go` (A Linha de Produção / Regras de Negócio)
Este componente aplica as regras do sistema (como não aceitar nomes vazios) e gera identificadores únicos (UUID) antes de salvar o registro.

```go
package service

import (
	"app/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CriarUsuarioService(nome, email, senha string) (*UserResponse, error) {
	// Regra de Negócio: Valida se o cliente preencheu o nome
	if nome == "" {
		return nil, errors.New("Nome não pode ser vazio")
	}

	// Gera uma string única universal (UUID) para servir como ID na tabela
	id := uuid.New().String()
	fmt.Println("ID gerado:", id)

	// Repassa os dados validados com o ID para serem salvos definitivamente no banco
	user, err := repository.SalvarAT_DB(id, nome, email, senha, time.Now())
	if err != nil {
		return nil, err
	}	

	return user, nil
}
```

---

### 4. `internal/repository/db.go` (O Gerador de Energia)
Responsável exclusivo por estabelecer e testar a ligação direta com a base de dados PostgreSQL.

```go
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver necessário para o Go falar com o Postgres
)

func OpenDB() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=usuario_lucas password=123456 dbname=banco_teste_lucas sslmode=disable"

	// Inicializa o pool de conexões técnicas
	db, err := sql.Open("postgres", dsn)

	// Dá um "toque" no banco para garantir que ele está respondendo de verdade
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco realizada com sucesso")
	return db, nil
}
```

---

### 5. `internal/repository/user_repository.go` (A Prensa / Inserção SQL)
Este arquivo executa o comando `INSERT` para gravar as informações nas colunas corretas da tabela.

```go
package repository

import (
	"app/internal/estrutura"
	"fmt"
	"time"
)

func SalvarAT_DB(id, nome, email, senha string, criado time.Time) (*estrutura.UserResponse, error) {
	// Abre a conexão com o banco de dados
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close() // Garante que a conexão será fechada ao terminar a função

	// Comando SQL puro para inserir os dados com segurança (\$1, \$2...) contra SQL Injection
	query := "INSERT INTO users (id, nome, email, senha, criado) VALUES (\$1, \$2, \$3, \$4, \$5)"
	_, ER := db.Exec(query, id, nome, email, senha, criado)
	if ER != nil {
		return nil, fmt.Errorf("Erro ao salvar usuário no banco de dados: %v", ER)
	}

	// Monta o objeto de resposta confirmando o que foi salvo
	user := &estrutura.UserResponse{
		ID:        id,
		Name:      nome,
		Email:     email,
		Senha:     senha,
		Criado_em: time.Now(),
	}

	return user, nil
}
```

---

### 6. `internal/estrutura/interface.go` (O Molde dos Dados)
Contém os modelos (structs) que moldam a entrada de requisições e a saída de respostas.

```go
package estrutura

import "time"

// UserResponse representa a estrutura completa devolvida ao cliente
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Senha     string    `json:"senha"`
	Email     string    `json:"email"`
	Criado_em time.Time `json:"criado_em"`
}

// UserRequest define quais campos são obrigatórios receber no momento da criação
type UserRequest struct {
	Name  string `json:"name"`
	Senha string `json:"senha"`
	Email string `json:"email"`
}	
```

---

## 🛠️ Dependências Utilizadas

Para que este módulo funcione corretamente, foram baixadas e instaladas as seguintes dependências externas via terminal:

```bash
# Roteador HTTP leve e veloz
go get github.com/go-chi/chi/v5

# Driver oficial do PostgreSQL para o Go
go get github.com/lib/pq

# Biblioteca para geração de IDs únicos universais
go get github.com/google/uuid
```

---

## 🚀 Como Testar a Inserção (via cURL)

# 🗄️ Estrutura do Banco de Dados para o Método POST

Para que o seu método **POST** consiga inserir os dados usando o código atual (com o comando `INSERT` nativo), o formato da tabela precisa bater exatamente com as colunas que você definiu na sua query SQL.

Aqui está o comando **`CREATE TABLE`** pronto para você executar no seu PostgreSQL.

---

### 💻 Comando SQL para Criar a Tabela

Abra o terminal do seu banco de dados (`banco_teste_lucas`) e execute o seguinte comando:

```sql
CREATE TABLE public.users (
    id character varying(36) NOT NULL,
    nome character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    senha character varying(100) NOT NULL,
    criado timestamp with time zone DEFAULT now(),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
```

> **Nota:** Use o código com cuidado ao executar em ambientes de produção.

---

### 🔍 Por que a tabela precisa ser exatamente assim?

Ao olhar o arquivo `banco.go` que você enviou, a sua linha de inserção é:  
`INSERT INTO users (id, nome, email, senha, criado) VALUES ($1, $2, $3, $4, $5)`

* **`id character varying(36)`**: O tamanho 36 é o tamanho exato de um **UUID** em formato de texto (que você está gerando no arquivo `service.go`). Ele foi definido como a chave primária (`PRIMARY KEY`).
* **`nome`, `email` e `senha`**: São mapeados como textos variáveis (`character varying`) com tamanhos seguros para guardar dados de cadastro.
* **`criado timestamp with time zone`**: Armazena a data e a hora exatas em que o usuário foi registrado. O `DEFAULT now()` garante que, se o Go esquecer de enviar a data por algum motivo, o próprio banco salva o horário atual automaticamente.

---

### 💡 Uma dica importante sobre as colunas

Note que na primeira etapa que fizemos juntos (o método GET), a sua tabela original possuía também a coluna `medicamento` e `data_nascimento`.

Como o seu código do **POST** atual envia apenas 5 campos (`id, nome, email, senha, criado`), criamos a tabela acima apenas com esses 5. Se você mantivesse os campos de medicamento e data de nascimento na tabela e eles fossem obrigatórios (`NOT NULL`), o seu POST daria erro por não enviá-los.

---

## 🚀 Próximos Passos

Crie a tabela no seu banco de dados e faça um teste rodando o seu **POST via cURL** ou Insomnia! 

<FollowUp>
Se o registro for inserido com sucesso, me avise se quer que eu te ajude a **adicionar os campos de medicamento e data de nascimento no POST** para que as duas partes fiquem idênticas!
</FollowUp>


Com o servidor rodando localmente, você pode disparar um comando no terminal utilizando a ferramenta **cURL** para simular um cliente inserindo dados na API:

```bash
curl -X POST http://localhost:8080/usuary \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Lucas Silva",
    "email": "lucas@email.com",
    "senha": "password123"
  }'
```

**Resposta esperada (JSON):**
```json
{
  "id": "e3b8a1c9-7d4f-4e3a-bc2d-89ef12345678",
  "name": "Lucas Silva",
  "senha": "password123",
  "email": "lucas@email.com",
  "criado_em": "2026-07-03T19:38:00Z"
}
```
