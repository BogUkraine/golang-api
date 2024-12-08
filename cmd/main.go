package main

import (
	"encoding/json"
	"log"
	controllers "main/internal/controllers"
	services "main/internal/services"
	"main/internal/storage"
	"net/http"

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

	var (
		userService    = services.NewUserService(storageInstance)
		userController = controllers.NewUserController(userService)
	)

	http.HandleFunc("GET /hello", helloHandler)
	http.HandleFunc("GET /users/{id}", userController.GetUserHandler)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
