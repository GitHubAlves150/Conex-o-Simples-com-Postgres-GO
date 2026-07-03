package handdlers

import (
	"api_teste/internal/estrutura"
	"encoding/json"
	"net/http"
)

//Este é o metodo GET tradicional que vocẽ já conhece
//Note o "(h *UserHanddler)" antes do nome da função.Isso dá acesso ao "h.db"

func (h *UserHandler) ObterUsuarios(w http.ResponseWriter, r *http.Request) {
	var users []estrutura.UserResponse

	//Usamos o banco atraves do h.db
	err := h.db.Find(&users).Error
	if err != nil {
		http.Error(w, "Erro ao conectar-se com o banco", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
