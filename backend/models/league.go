package models

//League will be for the league that will be using the ladder as a basic starting off point.
type League struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Game      string `json:"game"`
	Desc      string `json:"description"`
	Season    int    `json:"season"`
	Format    int    `json:"format"`
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
}
