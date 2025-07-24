package postgres

import (
	"auth-service/domain"
	"database/sql"
	"errors"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (id, username, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, to_timestamp($5), to_timestamp($6))
	`

	_, err := r.db.Exec(query,
		user.ID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.CreatedAt.Unix(),
		user.UpdatedAt.Unix(),
	)

	return err
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	row := r.db.QueryRow(query, email)

	var user domain.User
	var createdAt, updatedAt sql.NullTime

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Email bulunamadı
		}
		return nil, err
	}

	user.CreatedAt = createdAt.Time
	user.UpdatedAt = updatedAt.Time

	return &user, nil
}

func (r *PostgresUserRepository) FindByID(id string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var user domain.User
	var createdAt, updatedAt sql.NullTime

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user.CreatedAt = createdAt.Time
	user.UpdatedAt = updatedAt.Time

	return &user, nil
}
