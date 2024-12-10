package controller

import (
	"back-user-service/internal/user/entity"
	"back-user-service/internal/user/usecase"
	"back-user-service/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserController é o controlador para gerenciar os usuários
type UserController struct {
	usecase *usecase.UserUsecase
}

// NewUserController cria uma nova instância do controlador
func NewUserController(u *usecase.UserUsecase) *UserController {
	return &UserController{usecase: u}
}

// GetAllUsers retorna todos os usuários
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := uc.usecase.GetAllUsers()
	response.JSON(w, http.StatusOK, users)
}

// GetUserByID retorna um usuário pelo ID
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := uc.usecase.GetUserByID(id)
	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	response.JSON(w, http.StatusOK, user)
}

// CreateUser cria um novo usuário
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := uc.usecase.CreateUser(user)
	response.JSON(w, http.StatusCreated, createdUser)
}

// UpdateUser atualiza um usuário
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	updatedUser := uc.usecase.UpdateUser(id, user)
	if updatedUser == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	response.JSON(w, http.StatusOK, updatedUser)
}

// DeleteUser deleta um usuário
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if !uc.usecase.DeleteUser(id) {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
