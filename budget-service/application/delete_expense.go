package application

import (
	"budget-tracker/domain"
	"errors"
)

var ErrExpenseNotFound = errors.New("expense not found")

type DeleteExpenseCommand struct {
	ExpenseID string
	UserID    string // 👈 Eğer multi-tenant yapıdaysan kullanılır
}

type DeleteExpenseHandler struct {
	Repo domain.ExpenseRepository
}

func NewDeleteExpenseHandler(repo domain.ExpenseRepository) *DeleteExpenseHandler {
	return &DeleteExpenseHandler{Repo: repo}
}

func (h *DeleteExpenseHandler) Handle(cmd *DeleteExpenseCommand) error {
	// 1. Kayıt var mı kontrolü
	expense, err := h.Repo.GetByID(cmd.ExpenseID)
	if err != nil {
		return err // 500 olarak işlenebilir
	}
	if expense == nil {
		return ErrExpenseNotFound
	}

	// Eğer UserID kontrolü yapacaksan:
	if cmd.UserID != "" && cmd.UserID != expense.UserID {
		return ErrExpenseNotFound // aynı hatayı döner, bilgi sızdırmaz
	}

	// 2. Silme işlemi
	return h.Repo.Delete(cmd.ExpenseID)
}
