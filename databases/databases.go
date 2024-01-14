package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/KanhaiyaKumarGupta/assignment/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() error {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
        return err
    }

    dsn := os.Getenv("dsn") // Changed from "dsn" to "DATABASE_DSN"
    if dsn == "" {
        log.Fatalf("DATABASE_DSN not set in environment")
        return fmt.Errorf("DATABASE_DSN not set in environment")
    }
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
    if err != nil {
        return err
    }
    if err := db.AutoMigrate(&models.Task{}); err != nil {
        return err
    }
    sqlDB, err := db.DB()
    if err != nil {
        return err
    }

    if err = sqlDB.Ping(); err != nil {
        return err
    }

    fmt.Println("Connected to database successfully")
    return nil
}


func GetDB() *gorm.DB {
	return db
}
