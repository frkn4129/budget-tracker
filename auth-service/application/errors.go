package application

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email is already registered")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidCredentials = errors.New("invalid email or password")
)
