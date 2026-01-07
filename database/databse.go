package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Не найден .env файл, использую переменные окружения")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "ratingf1"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
	if password == "" {
		return fmt.Errorf("пароль БД не указан. Укажите в .env файле или переменных окружения")
	}

	targetConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err = sql.Open("postgres", targetConnStr)
	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err == nil {
			configureDB()
			log.Printf("Успешно подключено к базе данных '%s'!", dbname)

			if err := createTables(); err != nil {
				return fmt.Errorf("ошибка создания таблиц: %w", err)
			}

			return nil
		}
	}

	log.Printf("Не удалось подключиться к БД '%s'. Попытка создать базу данных...", dbname)

	serverConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=%s",
		host, port, user, password, sslmode,
	)

	tempDB, err := sql.Open("postgres", serverConnStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к серверу PostgreSQL: %w", err)
	}
	defer tempDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := tempDB.PingContext(ctx); err != nil {
		return fmt.Errorf("ошибка соединения с серверу PostgreSQL: %w", err)
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbname)
	err = tempDB.QueryRowContext(ctx, query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("ошибка проверки существования БД: %w", err)
	}

	if !exists {
		log.Printf("Создание базы данных '%s'...", dbname)

		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbname)
		_, err = tempDB.ExecContext(ctx, createQuery)
		if err != nil {
			return fmt.Errorf("ошибка создания базы данных: %w", err)
		}

		log.Printf("База данных '%s' успешно создана", dbname)
	} else {
		log.Printf("База данных '%s' уже существует", dbname)
	}

	db, err = sql.Open("postgres", targetConnStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к новой БД: %w", err)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	if err := db.PingContext(ctx2); err != nil {
		return fmt.Errorf("ошибка соединения с новой БД: %w", err)
	}

	configureDB()

	if err := createTables(); err != nil {
		return fmt.Errorf("ошибка создания таблиц: %w", err)
	}

	log.Printf("Успешно подключено к базе данных '%s'!", dbname)
	return nil
}

func configureDB() {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
}

func createTables() error {
	tables := []struct {
		name string
		sql  string
	}{
		{
			name: "racers",
			sql: `CREATE TABLE IF NOT EXISTS racers (
				id SERIAL PRIMARY KEY,
				country VARCHAR(100) NOT NULL,
				nameracers VARCHAR(100) NOT NULL,
				lastnameracers VARCHAR(100) NOT NULL,
				driveteam VARCHAR(100) NOT NULL
			)`,
		},
		{
			name: "teams",
			sql: `CREATE TABLE IF NOT EXISTS teams (
				id SERIAL PRIMARY KEY,
				nameteam VARCHAR(100) NOT NULL
			)`,
		},
		{
			name: "topracerc",
			sql: `CREATE TABLE IF NOT EXISTS topracerc (
				id SERIAL PRIMARY KEY,
				teamracers VARCHAR(100) NOT NULL,
				nameracer VARCHAR(100) NOT NULL,
				lastnameracer VARCHAR(100) NOT NULL,
				points VARCHAR(50) NOT NULL
			)`,
		},
		{
			name: "highway",
			sql: `CREATE TABLE IF NOT EXISTS highway (
				id SERIAL PRIMARY KEY,
				namehighway VARCHAR(100) NOT NULL,
				countryhighway VARCHAR(100) NOT NULL,
				lenght INTEGER NOT NULL
			)`,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, table := range tables {
		log.Printf("Создание/проверка таблицы '%s'...", table.name)

		_, err := db.ExecContext(ctx, table.sql)
		if err != nil {
			return fmt.Errorf("ошибка создания таблицы %s: %w", table.name, err)
		}

		log.Printf("Таблица '%s' готова", table.name)
	}

	log.Println("Все таблицы успешно созданы/проверены")
	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("База данных не инициализирована. Вызовите InitDatabase() сначала")
	}
	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
