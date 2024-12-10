package main

import (
	"back-user-service/internal/user/controller"
	"back-user-service/internal/user/repository"
	"back-user-service/internal/user/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Repositório
	userRepo := repository.NewUserRepository()

	// Casos de uso
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Controlador
	userController := controller.NewUserController(userUsecase)

	// Roteador
	router := mux.NewRouter()
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Inicia o servidor
	log.Println("Iniciando servidor na porta 8080...")
	http.ListenAndServe(":8080", router)
}
