package handdler

import (
	"app/internal/interface_test"
	"app/internal/service"
	"encoding/json"
	"net/http"
)

func CriaUsuarioHanddler(w http.ResponseWriter, r *http.Request) {
	var Req interface_test.UserResponse

	err := json.NewDecoder(r.Body).Decode(&Req)

	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := service.CriarUsuarioService(Req.Name, Req.Senha, Req.Email)
	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
