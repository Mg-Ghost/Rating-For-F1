package models

type Racers struct {
	ID int `json:"id"`
	Country string `json:"country"`
	Nameracers string `json:"nameracerf1"`
	Lastnameracers string `json:"lastnameracerf1"`
	Driveteam string `json:"drivetimef1"`
}

type Teams struct {
	ID int `json:"id"`
	Nameteam string `jsong:"nameteamsf1"`
}

type Topracerc struct {
	ID int `json:"id"`
	Teamracers string `json:"teamracersf1"`
	Nameracer string `json:"nameracer"`
	Lastnameracer string `json:"lastnameracer"`
	Points string `json:"points"`
}

type Highway struct {
	ID int `json:"id"`
	Namehighway string `json:"namehighway"`
	Countryhighway string `json:"countryhighway"`
	Lenght int `json:"lenght"`
}