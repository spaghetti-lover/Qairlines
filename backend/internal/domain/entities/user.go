package entities

import (
	"errors"
	"net/mail"

	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(name string, address string) (User, error) {

	email, err := validMailAddress(address)
	if err != nil {
		return User{}, errors.New("invalid email address")
	}

	uuid := utils.RandomString(10)

	user := User{
		ID:    uuid,
		Name:  name,
		Email: email,
	}

	return user, nil
}

func validMailAddress(address string) (string, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", err
	}

	return addr.Address, nil
}
