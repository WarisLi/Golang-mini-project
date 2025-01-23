package core

// secondary port
type UserRepository interface {
	Create(user User) error
	ValidateUser(requestUser User) error
}
