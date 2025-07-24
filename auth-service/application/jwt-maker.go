package application

import "time"

type JWTMaker interface {
	GenerateToken(userID string, duration time.Duration) (string, error)
}
