package highway

import (
	"RatingForF1/database"
	"RatingForF1/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ReadHighway(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Method not allowed"})
		return
	}

	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM highway")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Databse error" + err.Error()})
		return
	}
	defer rows.Close()

	HighwayList := make([]models.Highway, 0)
	for rows.Next() {
		var highway models.Highway
		err := rows.Scan(&highway.ID, &highway.Namehighway, &highway.Countryhighway, &highway.Lenght)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error reading rows" + err.Error()})
			return
		}
		HighwayList = append(HighwayList, highway)
	}
	json.NewEncoder(w).Encode(HighwayList)
}

func GetHighwayWrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset = utf-8")

	if r.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed!"})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/highway/")
	if path == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalid requred"})
		return
	}

	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	
	GetHighwayById(w, idInt)
}

func GetHighwayById(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json: charset = utf-8")
	db := database.GetDB()

	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}

	row := db.QueryRow("SELECT * FROM highway WHERE id = $1", id)
	highway := models.Highway{}
	err := row.Scan(&highway.ID, &highway.Namehighway, &highway.Countryhighway, &highway.Lenght)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Highway not found"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(highway)
}
