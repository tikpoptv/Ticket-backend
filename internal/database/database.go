package database

import (
	"fmt"
	"ticket-backend/internal/config"
	"ticket-backend/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database",
			logger.String("error", err.Error()),
			logger.String("host", cfg.DBHost),
			logger.String("port", cfg.DBPort),
			logger.String("database", cfg.DBName),
		)
		return err
	}
	DB = db
	logger.Info("Successfully connected to database",
		logger.String("host", cfg.DBHost),
		logger.String("port", cfg.DBPort),
		logger.String("database", cfg.DBName),
	)
	return nil
}
