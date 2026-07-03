curl -X POST http://localhost:8080/Perfil \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com",
    "senha": "123456",
    "medicamento": "Losartana 50mg",
    "data_nascimento": "1990-01-15"
  }'



  Nome string `json:"nome"`
    Email string `json:"email"`
    Senha string `json:"senha"`
    Medicamento string `json:"medicamento"`
    Data_nascimento string `json:"data_nascimento"`