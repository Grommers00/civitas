package models

//Match will be a match between two competitors where we can have a winner, tie, and loser.
type Match struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Game      string `json:"game"`
	Map       string `json:"map"`
	Desc      string `json:"description"`
	Season    int    `json:"season"`
	Format    int    `json:"format"`
	Score1    int    `json:"score"`
	Score2    int    `json:"score2"`
	Player1   int    `json:"player1"`
	Player1Id int    `json: "player1id"`
	Player2   int    `json:"player2"`
	Player2Id int    `json: "player2id"`
}
