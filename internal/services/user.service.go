package services

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	"net/http"
	"os"
	"time"
)

type UserStorage interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
}

type UserService struct {
	UserStorage UserStorage
}

// func NewUserService(userStorage UserStorage) *UserService {
// 	return &UserService{userStorage}
// }

func (userService *UserService) CreateUser(user *models.User) error {
	return userService.UserStorage.CreateUser(user)
}

func (userService *UserService) GetUser(id int) (*models.User, error) {
	user, err := userService.UserStorage.GetUser(id)
	return user, err
}

const charityMCC = 8398

func (userService *UserService) GetUserDonations(date time.Time) ([]models.Donation, error) {
	accountId := os.Getenv("MONO_ACCOUNT_ID") // should be obtained from https://api.monobank.ua/personal/client-info
	token := os.Getenv("MONO_API_KEY")

	if accountId == "" || token == "" {
		return nil, fmt.Errorf("MONO_* secret is not set")
	}

	url := fmt.Sprintf("https://api.monobank.ua/personal/statement/%v/%v/%v", accountId, date.Unix(), date.Add(24*time.Hour).Unix())

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Token", token)

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Default().Print(resp)
		return nil, fmt.Errorf("failed to get donations: status code %d", resp.StatusCode)
	}

	var donations []models.Donation
	if err := json.NewDecoder(resp.Body).Decode(&donations); err != nil {
		return nil, err
	}

	var filteredDonations []models.Donation
	for _, donation := range donations {
		if donation.MCC == charityMCC {
			filteredDonations = append(filteredDonations, donation)
		}
	}
	return filteredDonations, nil
}
