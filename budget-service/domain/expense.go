// domain/expense.go
package domain

type Expense struct {
	ID          string
	UserID      string
	Amount      float64
	Category    string
	Description string
	CreatedAt   int64 // Unix timestamp
}

type ExpenseRepository interface {
	Save(expense *Expense) error
	FindByID(id string) (*Expense, error)
	FindByUser(userID string) ([]*Expense, error)
	GetByID(id string) (*Expense, error)
	Delete(id string) error
}
