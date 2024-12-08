package controllers

import (
	"encoding/json"
	"main/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// type UserService interface {
// 	GetUser(id int) (*models.User, error)
// 	CreateUser(user *models.User) error
// }

func (userController *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func (userController *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	print(vars)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := userController.userService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
