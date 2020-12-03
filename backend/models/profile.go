package models

//Profile will be for the users that will be using the ladder as a basic starting off point.
type Profile struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Race    string `json:"race"`
	Desc    string `json:"desc"`
	Status  string `json:"status"`
	Wins    int    `json:"wins"`
	Matches int    `json:"matches"`
	Games   string `json:"games"`
}
