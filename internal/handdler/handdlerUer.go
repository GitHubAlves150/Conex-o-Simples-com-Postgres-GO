package handdler

import (
	"app/internal/interface_test"
	"app/internal/service"
	"encoding/json"
	"net/http"
)

type UsuarioHandler struct {
	srv service.UsuarioService
}

func NewUsuarioHandler(srv service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{srv: srv}
}

func (h *UsuarioHandler) CriaUsuarioHanddler(w http.ResponseWriter, r *http.Request) {
	var Req interface_test.UserRequest

	err := json.NewDecoder(r.Body).Decode(&Req)
	if err != nil {
		http.Error(w, "Erro no payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.srv.CriarUsuario(Req.Name, Req.Email, Req.Senha)
	if err != nil {
		http.Error(w, "Erro no service: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}