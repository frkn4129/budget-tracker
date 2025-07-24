package application

import (
	"budget-tracker/domain"
	"time"

	"github.com/google/uuid"
)

// Dış dünyadan gelen veri değil
// Uygulama içi "bir işin yapılmasını" temsil eder
// DTO gibi de düşün
type CreateExpenseCommand struct {
	UserID      string
	Amount      float64
	Category    string
	Description string
}

// Response tipini sade tuttuk
type CreateExpenseResponse struct {
	ID string
}

type CreateExpenseHandler struct {
	Repo domain.ExpenseRepository
}

func NewCreateExpenseHandler(repo domain.ExpenseRepository) *CreateExpenseHandler {
	return &CreateExpenseHandler{Repo: repo}
}

func (h *CreateExpenseHandler) Handle(cmd *CreateExpenseCommand) (*CreateExpenseResponse, error) {
	expense := &domain.Expense{
		ID:          uuid.New().String(),
		UserID:      cmd.UserID,
		Amount:      cmd.Amount,
		Category:    cmd.Category,
		Description: cmd.Description,
		CreatedAt:   time.Now().Unix(), // Unix timestamp
	}

	if err := h.Repo.Save(expense); err != nil {
		return nil, err
	}

	return &CreateExpenseResponse{ID: expense.ID}, nil
}
