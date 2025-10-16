package highway

import "database/sql"

type Highway struct {
	ID int `json:"id"`
	Namehighway string `json:"namehighway"`
	Countryhighway string `json:"countryhighway"`
	Lenght int `json:"lenght"`
}

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}