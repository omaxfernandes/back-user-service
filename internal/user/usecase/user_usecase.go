package usecase

import (
	"back-user-service/internal/user/entity"
	"back-user-service/internal/user/repository"
)

// UserUsecase define os métodos de negócio para o usuário
type UserUsecase struct {
	repo *repository.UserRepository
}

// NewUserUsecase cria uma nova instância de UserUsecase
func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// GetAllUsers retorna todos os usuários
func (u *UserUsecase) GetAllUsers() []entity.User {
	return u.repo.GetAllUsers()
}

// GetUserByID retorna um usuário pelo ID
func (u *UserUsecase) GetUserByID(id int) *entity.User {
	return u.repo.GetUserByID(id)
}

// CreateUser cria um novo usuário
func (u *UserUsecase) CreateUser(user entity.User) entity.User {
	return u.repo.CreateUser(user)
}

// UpdateUser atualiza um usuário existente
func (u *UserUsecase) UpdateUser(id int, user entity.User) *entity.User {
	return u.repo.UpdateUser(id, user)
}

// DeleteUser deleta um usuário pelo ID
func (u *UserUsecase) DeleteUser(id int) bool {
	return u.repo.DeleteUser(id)
}
