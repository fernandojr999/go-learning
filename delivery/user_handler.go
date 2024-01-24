// delivery/user_handler.go
package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/fernandojr999/go-learning/domain"
	"github.com/fernandojr999/go-learning/usecase"
)

// UserHandler lida com as solicitações relacionadas a usuários
type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

// NewUserHandler cria uma instância de UserHandler
func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// CreateUser manipula solicitações de criação de usuário
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userUsecase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AuthenticateUser manipula solicitações de autenticação de usuário
func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var inputUser domain.User
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userUsecase.AuthenticateUser(&inputUser)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Autenticação bem-sucedida
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Autenticação bem-sucedida"))
}
