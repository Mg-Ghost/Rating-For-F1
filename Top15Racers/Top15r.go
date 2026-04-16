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
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	db := database.GetDB()

	rows, err := db.Query("SELECT id, teamracers, nameracer, lastnameracer, points FROM topracerc LIMIT 15")
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	defer rows.Close()

	topList := make([]models.Topracerc, 0, 15)

	for rows.Next() {
		var top models.Topracerc
		if err := rows.Scan(&top.ID, &top.Teamracers, &top.Nameracer, &top.Lastnameracer, &top.Points); err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error reading row: "+err.Error())
			return
		}
		topList = append(topList, top)
	}

	writeJSON(w, http.StatusOK, topList)
}

func GetTopRacersWrapper(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/Topracers/")
	if path == "" {
		writeJSONError(w, http.StatusBadRequest, "ID is required")
		return
	}

	idInt, err := strconv.Atoi(path)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	GetTopRacersById(w, idInt)
}

func GetTopRacersById(w http.ResponseWriter, id int) {
	db := database.GetDB()
	if db == nil {
		writeJSONError(w, http.StatusInternalServerError, "Database not initialized")
		return
	}

	// FIX: теперь реально используем id
	row := db.QueryRow("SELECT id, teamracers, nameracer, lastnameracer, points FROM topracerc WHERE id = ?", id)

	var top models.Topracerc
	err := row.Scan(&top.ID, &top.Teamracers, &top.Nameracer, &top.Lastnameracer, &top.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			writeJSONError(w, http.StatusNotFound, "Racer not found")
		} else {
			writeJSONError(w, http.StatusInternalServerError, "Database error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, top)
}

// --- helpers (не ломают архитектуру, но убирают дублирование) ---

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
