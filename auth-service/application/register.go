package application

import (
	"auth-service/domain"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	Username string
	Email    string
	Password string
}

type RegisterResponse struct {
	ID       string
	Username string
	Email    string
}

type RegisterHandler struct {
	Repo domain.UserRepository
}

func NewRegisterHandler(repo domain.UserRepository) *RegisterHandler {
	return &RegisterHandler{Repo: repo}
}

func (h *RegisterHandler) Handle(cmd *RegisterCommand) (*RegisterResponse, error) {
	// Email kontrolü
	existing, err := h.Repo.FindByEmail(cmd.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	// Şifre hashleme
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           uuid.New().String(),
		Username:     cmd.Username,
		Email:        cmd.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := h.Repo.Create(user); err != nil {
		return nil, err
	}

	return &RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
