package main

import (
	"encoding/json"
	"log"
	"main/internal/controllers"
	"main/internal/services"
	"main/internal/storage"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	storageInstance, err := storage.NewStorageInstance()
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	userService := services.UserService{UserStorage: storageInstance}
	userController := controllers.UserController{UserService: &userService}

	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	// router.HandleFunc("/users", userController.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/me/donations", userController.GetUserDonationsHandler).Methods("GET")

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
