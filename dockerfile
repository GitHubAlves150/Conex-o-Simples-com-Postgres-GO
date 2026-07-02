# Estágio 1: Compilar o código Go
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar os ficheiros de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte e compilar
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Estágio 2: Executar o binário num ambiente limpo e leve
FROM alpine:latest

WORKDIR /app

# Copiar o binário gerado no estágio anterior
COPY --from=builder /app/main .

# Porta que a sua API Go escuta (ajuste se não for a 8080)
EXPOSE 8080

CMD ["cmd/api/main"]
