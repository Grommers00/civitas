package models

// Season is the high level container
// which is referenced by each Standing structure
type Season struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Game  string `json:"game"`
}

// Standing is the table entry which references
// a player's performance in a season
type Standing struct {
	ID       int `json:"id"`
	SeasonID int `json:"seasonId"`
	PlayerID int `json:"playerId"`
	Wins     int `json:"wins"`
	Loses    int `json:"loses"`
	Matches  int `json:"matches"`
}
