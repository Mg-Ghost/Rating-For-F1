package racers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ReadRacers(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}

	if r.Method != "GET"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	rows, err := db.Query("SELECT * FROM racersf1")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error " + err.Error()})
		return
	}
	defer rows.Close()

	racersList := make([]Racers, 0)
	for rows.Next() {
		var racers Racers
		err := rows.Scan(&racers.ID, &racers.Country, &racers.Nameracers, &racers.Lastnameracers, &racers.Driveteam)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error raeding row " + err.Error()})
			return
		}
		racersList = append(racersList, racers)
	}
	json.NewEncoder(w).Encode(racersList)
}

func GetRacersWrapper(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset = utf-8")

	if r.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/racers/")
	if path == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalid requred"})
		return
	}

	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	GetRacersById(w, idInt)
}

func GetRacersById(w http.ResponseWriter, id int){
	w.Header().Set("Content-Type", "application/json; charset = utf-8")

	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}

	row := db.QueryRow("SELECT *  FROM racersf1 WHERE id = $1", id)

	racers := Racers{}
	err := row.Scan(&racers.ID, &racers.Country, &racers.Nameracers, &racers.Lastnameracers, &racers.Driveteam)
		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(w).Encode(map[string]string{"error": "Progress not found"}) 
			} else {
				json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
			}
			return
		}
	
	json.NewEncoder(w).Encode(racers)
}