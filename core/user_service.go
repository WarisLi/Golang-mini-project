package core

import (
	"golang.org/x/crypto/bcrypt"
)

// primary port
type UserService interface {
	RegisterUser(user User) error
	LoginUser(user User) error
}

// connect secondary port
type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// business logic
func (s *userServiceImpl) RegisterUser(user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// call secondary port
	err = s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) LoginUser(user User) error {
	err := s.repo.ValidateUser(user)
	if err != nil {
		return err
	}

	return nil
}
