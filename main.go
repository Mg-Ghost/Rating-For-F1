package main

import (
	"fmt"
	"log"
	"net/http"

	highway "RatingForF1/Highway"
	racers "RatingForF1/Racers"
	teamsf1 "RatingForF1/TeamsF1"
	top15racers "RatingForF1/Top15Racers"
	"RatingForF1/database"

	_ "github.com/lib/pq"
)

func main() {
	if err := database.InitDatabase(); err != nil {
		log.Fatal("Ошибка инициализации БД: ", err)
	}
	defer database.Close()

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
			teamsf1.ReadTeamF1(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/teams/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			teamsf1.GetTeamsWrapper(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//Маршруты для Highway
	http.HandleFunc("/highway", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			highway.ReadHighway(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/highway/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			highway.GetHighwayWrapper(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8182")
	log.Fatal(http.ListenAndServe(":8182", nil))
}
