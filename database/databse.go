package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDatabase() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Ошибка загрузки .env файла: %w", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if password == "" {
		return fmt.Errorf("Пароль БД не указан в .env файле")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Ошибка подключение к БД: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("Ошибка соединения: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Успешно подключено к PostgreSQL!")
	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("База данных не инициализирована")
	}
	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
