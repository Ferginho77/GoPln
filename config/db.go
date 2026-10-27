package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
 	"gorm.io/driver/mysql"
 	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	_ = godotenv.Load()

	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	password := os.Getenv("DB_PASSWORD")
	database := getEnv("DB_NAME", "plndb")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to initialize database pool: ", err)
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	fmt.Printf("Database connection successful: %s/%s\n", host, database)
 }

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}