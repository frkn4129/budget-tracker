package infrastructure

import (
	"budget-tracker/domain"
	"database/sql"
)

type PostgresExpenseRepository struct {
	db *sql.DB
}

func NewPostgresExpenseRepository(db *sql.DB) *PostgresExpenseRepository {
	return &PostgresExpenseRepository{db: db}
}

func (r *PostgresExpenseRepository) Save(expense *domain.Expense) error {
	query := `
		INSERT INTO expenses (id, user_id, amount, category, description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query,
		expense.ID,
		expense.UserID,
		expense.Amount,
		expense.Category,
		expense.Description,
		expense.CreatedAt,
	)
	return err
}

func (r *PostgresExpenseRepository) FindByID(id string) (*domain.Expense, error) {
	query := `SELECT id, user_id, amount, category, description, created_at FROM expenses WHERE id = $1`

	row := r.db.QueryRow(query, id)
	var e domain.Expense
	err := row.Scan(&e.ID, &e.UserID, &e.Amount, &e.Category, &e.Description, &e.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found, üst katman bunu kontrol eder
		}
		return nil, err
	}
	return &e, nil
}

func (r *PostgresExpenseRepository) FindByUser(userID string) ([]*domain.Expense, error) {
	query := `SELECT id, user_id, amount, category, description, created_at FROM expenses WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*domain.Expense
	for rows.Next() {
		var e domain.Expense
		if err := rows.Scan(&e.ID, &e.UserID, &e.Amount, &e.Category, &e.Description, &e.CreatedAt); err != nil {
			return nil, err
		}
		expenses = append(expenses, &e)
	}
	return expenses, nil
}

// ✅ Yeni: Delete fonksiyonu
func (r *PostgresExpenseRepository) Delete(id string) error {
	query := `DELETE FROM expenses WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresExpenseRepository) GetByID(id string) (*domain.Expense, error) {
	return r.FindByID(id)
}
