package repository

import (
	"back-user-service/internal/user/entity"
	"sync"
)

var users = []entity.User{
	{ID: 1, Name: "João Silva", Email: "joao@gmail.com"},
	{ID: 2, Name: "Maria Souza", Email: "maria@gmail.com"},
}

var mu sync.Mutex

// UserRepository é uma estrutura que implementa os métodos de persistência de usuários
type UserRepository struct{}

// NewUserRepository cria uma nova instância do repositório
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetAllUsers retorna todos os usuários
func (r *UserRepository) GetAllUsers() []entity.User {
	return users
}

// GetUserByID retorna um usuário pelo ID
func (r *UserRepository) GetUserByID(id int) *entity.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

// CreateUser adiciona um novo usuário
func (r *UserRepository) CreateUser(user entity.User) entity.User {
	mu.Lock()
	defer mu.Unlock()
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

// UpdateUser atualiza um usuário existente
func (r *UserRepository) UpdateUser(id int, user entity.User) *entity.User {
	for i, u := range users {
		if u.ID == id {
			users[i] = user
			return &users[i]
		}
	}
	return nil
}

// DeleteUser remove um usuário
func (r *UserRepository) DeleteUser(id int) bool {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
