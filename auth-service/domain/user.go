package domain

import "time"

// User entity'si: veritabanı şemasıyla birebir değil, domain odaklıdır.
type User struct {
	ID           string // UUID
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// UserRepository: dış dünyaya (PostgreSQL, Redis, test, vs.) karşı soyutlama
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}
