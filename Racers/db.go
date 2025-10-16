package racers

import "database/sql"

type Racers struct {
	ID int `json:"id"`
	Country string `json:"country"`
	Nameracers string `json:"nameracerf1"`
	Lastnameracers string `json:"lastnameracerf1"`
	Driveteam string `json:"drivetimef1"`
}

var db *sql.DB

func InitDB(database *sql.DB){
	db = database
}
