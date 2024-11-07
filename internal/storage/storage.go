package Storage

import (
	"fmt"
	"github.com/Art4mPanin/gRPCAuthService/internal/config"
	"github.com/Art4mPanin/gRPCAuthService/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	pass := os.Getenv("POSTGRES_PASSWORD")

	if pass == "" {
		log.Fatalf("POSTGRES_PASSWORD not set in .env file")
	}

	fmt.Printf("DB Config: Host=%s, Port=%d, Username=%s, Database=%s\n",
		cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.Username, cfg.DBConfig.Database)

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable password=%s",
		cfg.DBConfig.Host, cfg.DBConfig.Username, cfg.DBConfig.Database, cfg.DBConfig.Port, pass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Roles{}, &models.UserRoles{}, &models.Info{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
