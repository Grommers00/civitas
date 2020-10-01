package models

//Profile will be for the users that will be using the ladder as a basic starting off point.
type Profile struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}
