package interface

package interface

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Senha string `json:"senha"`
	Email string `json:"email"`
	Criado_em time.Time `json:"criado_em"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Senha string `json:"senha"`
	Email string `json:"email"`
}	