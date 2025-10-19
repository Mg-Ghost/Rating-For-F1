package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"RatingForF1/Highway"
	"RatingForF1/Racers"
	"RatingForF1/TeamsF1"
	"RatingForF1/Top15Racers"

	_ "github.com/lib/pq"
)

var db *sql.DB

func readPasswordFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func main() {
    password, err := readPasswordFromFile("pass.txt")
    if err != nil {
        log.Fatal("Ошибка чтения пароля!")
    }

    connStr := fmt.Sprintf("user = postgres password=%s dbname = ratingf1 sslmode = disable", password)

    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Ошибка подключения к БД!")
    }

    defer db.Close()

    err = db.Ping()
    if err != nil{
        log.Fatal("Ошибка Ping:", err)
    }

    fmt.Println("Успешно подключено к PostgreSQL (база RatingF1)!")

    highway.InitDB(db)
    racers.InitDB(db)
    teamsf1.InitDB(db)
    top15racers.InitDB(db)

// Маршруты для Racers
	http.HandleFunc("/racers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			racers.ReadRacers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/racers/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        racers.GetRacersWrapper(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
	})

// Маршруты для Топ 15
		http.HandleFunc("/Topracers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			top15racers.ReadTopRacers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/Topracers/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        top15racers.GetTopRacersWrapper(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
	})

// Маршруты для Teams
	http.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			teamsf1.ReadTeamF1(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/teams/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			teamsf1.GetTeamsWrapper(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8182")
	log.Fatal(http.ListenAndServe(":8182", nil))
}