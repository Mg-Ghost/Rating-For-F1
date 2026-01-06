package Top15Racers

import (
	"RatingForF1/database"
	"RatingForF1/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ReadTopRacers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM rating_top_15 LIMIT 15")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error " + err.Error()})
		return
	}
	defer rows.Close()

	TopList := make([]models.Topracerc, 0)
	for rows.Next() {
		var Top models.Topracerc
		err := rows.Scan(&Top.ID, &Top.Teamracers, &Top.Nameracer, &Top.Lastnameracer, &Top.Points)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error raeding row " + err.Error()})
			return
		}
		TopList = append(TopList, Top)
	}
	json.NewEncoder(w).Encode(TopList)
}

func GetTopRacersWrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset = utf-8")

	if r.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/Topracers/")
	if path == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalid requred"})
		return
	}

	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	GetTopRacersById(w, idInt)
}

func GetTopRacersById(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json; charset = utf-8")
	db := database.GetDB()

	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}

	row := db.QueryRow("SELECT * FROM rating_top_15 ORDER BY id LIMIT 15;", id)

	Top := models.Topracerc{}
	err := row.Scan(&Top.ID, &Top.Teamracers, &Top.Nameracer, &Top.Lastnameracer, &Top.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Racers not found"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		}
		return
	}

	json.NewEncoder(w).Encode(Top)
}
