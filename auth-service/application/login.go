package application

import (
	"auth-service/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

type LoginHandler struct {
	Repo     domain.UserRepository
	JWTMaker JWTMaker
}

func NewLoginHandler(repo domain.UserRepository, jwtMaker JWTMaker) *LoginHandler {
	return &LoginHandler{
		Repo:     repo,
		JWTMaker: jwtMaker,
	}
}

func (h *LoginHandler) Handle(cmd *LoginCommand) (*LoginResponse, error) {
	user, err := h.Repo.FindByEmail(cmd.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cmd.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := h.JWTMaker.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token}, nil
}
