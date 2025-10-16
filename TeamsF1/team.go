package teamsf1

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ReadTeamF1(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		json.NewEncoder(w).Encode((map[string]string{"error": "Database not initialized"}))
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error":"Method not allowed"})
		return
	}

	rows,err := db.Query("select * from teamsf1")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error" + err.Error()})
		return
	}
	defer rows.Close()

	teamsList := make([]Teams, 0)
	for rows.Next() {
		var teams Teams
		err := rows.Scan(&teams.ID, &teams.Nameteam)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Error reading rows" + err.Error()})
		}
		teamsList = append(teamsList, teams)
	}
	json.NewEncoder(w).Encode(teamsList)
}

func GetTeamsWrapper(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset = utf-8")

	if r.Method != "GET"{
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/teams/")
	if path == ""{
		json.NewEncoder(w).Encode(map[string]string{"error": "ID Invalid requred"})
		return
	}

	idInt, err := strconv.Atoi(path)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID Invalid"})
		return
	}

	GetTeamsID(w,idInt)
}

func GetTeamsID(w http.ResponseWriter, id int){
	w.Header().Set("Content-Type", "application/json; charset = utf-8")

	if db == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Database not initialized"})
		return
	}

	row := db.QueryRow("SELECT * FROM teamsf1 WHERE id = $1", id)

	teams := Teams{}
	err := row.Scan(&teams.ID, &teams.Nameteam)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Teams not found"}) 
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Database error: " + err.Error()})
		}
		return
	}
	
json.NewEncoder(w).Encode(teams)
}