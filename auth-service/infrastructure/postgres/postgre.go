package postgres

import (
	"auth-service/internal/config"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	// Port integer ise string'e dönüştür
	portStr := strconv.Itoa(cfg.DBPort)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		portStr,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}

	return db, nil
}
