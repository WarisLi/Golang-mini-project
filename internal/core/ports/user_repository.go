package ports

import (
	"github.com/WarisLi/Golang-mini-project/internal/core/models"
)

// secondary port
type UserRepository interface {
	Create(user models.User) error
	ValidateUser(requestUser models.User) error
}
