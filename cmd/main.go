package main

import (
	"fmt"
	"net/http"
	"ticket-backend/internal/config"
	"ticket-backend/internal/database"
	"ticket-backend/internal/logger"
)

func main() {
	logger.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration",
			logger.String("error", err.Error()))
	}

	if err := database.ConnectDB(cfg); err != nil {
		logger.Fatal("Failed to connect to the database",
			logger.String("error", err.Error()))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API!"))
	})

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	logger.Info("Server starting",
		logger.String("port", cfg.ServerPort))

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		logger.Fatal("Server failed to start",
			logger.String("error", err.Error()))
	}
}
