package models

//News will be for the articles that will be written regarding the league
type News struct {
	ID      int    `json:"ID"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Article string `json:"article"`
	Date    string `json:"date"`
}
