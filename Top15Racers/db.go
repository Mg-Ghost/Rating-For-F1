package top15racers

import "database/sql"

type Topracerc struct {
	ID int `json:"id"`
	Teamracers string `json:"teamracersf1"`
	Nameracer string `json:"nameracer"`
	Lastnameracer string `json:"lastnameracer"`
	Points string `json:"points"`
}

var db *sql.DB

func InitDB(database *sql.DB){
	db = database
}