package models

/*We define the models for the json data.
Comments is a subset of the Game data so has its own struct
Games struct contains an array of games.
*/

type Games struct {
	Games []Game `json:"games"`
}

type Comment struct {
	User        string `json:"user"`
	Message     string `json:"message"`
	DateCreated string `json:"dateCreated"`
	Like        int    `json:"like"`
}

type Game struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	By          string  `json:"by"`
	Platform    string  `json:"platform"`
	AgeRating   string  `json:"age_rating"`
	Likes       int     `json:"likes"`
	Comment     Comment `json:"comments"`
}
