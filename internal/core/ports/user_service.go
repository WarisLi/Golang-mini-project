package ports

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(usernamePassword models.UsernamePassword) error
	LoginUser(usernamePassword models.UsernamePassword) (string, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

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

func (s *userServiceImpl) LoginUser(requestUser models.UsernamePassword) (string, error) {
	userData, err := s.repo.GetUser(requestUser.Username)
	if err != nil {
		return "", err
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(requestUser.Password))
	if err != nil {
		return "", err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": userData.Username,
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
