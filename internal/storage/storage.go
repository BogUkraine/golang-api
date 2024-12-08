package storage

import (
	"database/sql"
	"log"
	"main/internal/models"
	"os"

	_ "github.com/lib/pq"
)

type StorageInstance struct {
	db *sql.DB
}

func NewStorageInstance() (*StorageInstance, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("One or more required environment variables are not set")
	}
	connectionString := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to database")

	return &StorageInstance{db}, nil
}

func (si *StorageInstance) CreateUser(user *models.User) error {
	err := si.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	return err
}

func (si *StorageInstance) GetUser(id int) (*models.User, error) {
	row := si.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("User not found: %v", err) // not fatal
			return nil, nil
		}

		log.Fatalf("Failed to get user: %v", err)
		return nil, err
	}

	return &user, nil
}
