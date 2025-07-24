package infrastructure

import (
	//"budget-tracker/domain"
	"budget-tracker/domain"
	"errors"
	"sync"
)

type InMemoryExpenseRepository struct {
	data map[string]*domain.Expense
	mu   sync.RWMutex
}

func NewInMemoryExpenseRepository() *InMemoryExpenseRepository {
	return &InMemoryExpenseRepository{
		data: make(map[string]*domain.Expense),
	}
}

func (r *InMemoryExpenseRepository) Save(expense *domain.Expense) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[expense.ID]; exists {
		return errors.New("expense already exists")
	}

	r.data[expense.ID] = expense
	return nil
}

func (r *InMemoryExpenseRepository) FindByID(id string) (*domain.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if exp, exists := r.data[id]; exists {
		return exp, nil
	}

	return nil, errors.New("expense not found")
}

func (r *InMemoryExpenseRepository) FindByUser(userID string) ([]*domain.Expense, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.Expense
	for _, e := range r.data {
		if e.UserID == userID {
			result = append(result, e)
		}
	}
	return result, nil
}
