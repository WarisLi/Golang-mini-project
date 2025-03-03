package ports

import (
	"github.com/WarisLi/Golang-mini-project/internal/core/models"
)

type UserRepository interface {
	GetUser(username string) (*models.User, error)
	Create(user models.User) error
}
