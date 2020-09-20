package models

//Users will be for the users that will be using the ladder as a basic starting off point.
type News struct {
	ID      int    `json:"ID"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Article string `json:"article"`
	Date    string `json:"date"`
}
