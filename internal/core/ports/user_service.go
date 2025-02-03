package ports

import (
	"encoding/json"
	"fmt"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"golang.org/x/crypto/bcrypt"
)

// primary port
type UserService interface {
	RegisterUser(usernamePassword models.UsernamePassword) error
	LoginUser(usernamePassword models.UsernamePassword) error
}

// connect secondary port
type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// business logic
func (s *userServiceImpl) RegisterUser(usernamePassword models.UsernamePassword) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usernamePassword.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usernamePassword.Password = string(hashedPassword)

	// Convert user to JSON
	data, err := json.Marshal(usernamePassword)
	if err != nil {
		fmt.Println("Error marshalling UsernamePassword:", err)
		return err
	}

	// Convert JSON to Product
	var user models.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error unmarshalling to User:", err)
		return err
	}

	// call secondary port
	err = s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) LoginUser(usernamePassword models.UsernamePassword) error {
	// Convert user to JSON
	data, err := json.Marshal(usernamePassword)
	if err != nil {
		fmt.Println("Error marshalling UsernamePassword:", err)
		return err
	}

	// Convert JSON to Product
	var user models.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error unmarshalling to User:", err)
		return err
	}

	err = s.repo.ValidateUser(user)
	if err != nil {
		return err
	}

	return nil
}
