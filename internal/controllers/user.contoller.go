package controllers

import (
	"encoding/json"
	"main/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserDonations(date time.Time) ([]models.Donation, error)
}

type UserController struct {
	UserService UserService
}

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

	user, err := userController.UserService.GetUser(id)
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

const TIME_FORMAT = "2006-01-02"

func (userController *UserController) GetUserDonationsHandler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		http.Error(w, "Date parameter is missing", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse(TIME_FORMAT, date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	dotations, err := userController.UserService.GetUserDonations(parsedDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dotations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
