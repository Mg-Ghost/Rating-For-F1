package teamsf1

import "database/sql"

type Teams struct {
	ID int `json:"id"`
	Nameteam string `jsong:"nameteamsf1"`
}

var db *sql.DB

func InitDB(database *sql.DB){
	db = database
}
