package entities

import "errors"

var (
	ErrEmailAlreadyUsed = errors.New("email already in use")
)
