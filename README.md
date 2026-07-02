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
Feito com ❤️ por [Seu Nome](https://github.com)
