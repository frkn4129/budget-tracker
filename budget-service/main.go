package main

import (
	"budget-tracker/application"
	"budget-tracker/infrastructure"
	"budget-tracker/interfaces/http"
	"budget-tracker/internal/config"
	"budget-tracker/internal/consul"
	"budget-tracker/internal/logger"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	logger.InitLogger(logger.NewZapLogger())
	cfg := config.Load()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("DB bağlantısı başarısız", "error", err)
	}
	defer db.Close()

	repo := infrastructure.NewPostgresExpenseRepository(db)
	createHandler := application.NewCreateExpenseHandler(repo)

	app := http.NewFiberServer(createHandler)

	port := ":3002"
	logger.Info("Fiber sunucu başlatıldı", "port", port)

	// Consul register
	portStr := strings.TrimPrefix(port, ":")
	p, _ := strconv.Atoi(portStr)
	go consul.RegisterToConsul("budget-service-"+portStr, "budget-service", "localhost", p)

	app.Listen(port)
}
