package models

import "time"

type JokeDTO struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Joke     string `json:"joke"`
}

type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Joke     string `json:"joke"`
	ID       int    `json:"id"`
}

type Joke struct {
	ID        int       `json:"id"`
	JokeID    string    `json:"joke_id"`
	Category  string    `json:"category"`
	Joke      string    `json:"joke"`
	CreatedAt time.Time `json:"created_at"`
}

type Cat struct {
	ID        int       `db:"id"`
	CatID     string    `db:"cat_id"`
	URL       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
}

type CatDTO struct {
	CatID string `db:"cat_id"`
	URL   string `db:"url"`
}

type NumberDTO struct {
	Number int    `json:"number"`
	Fact   string `json:"text"`
}

type Number struct {
	ID        int    `json:"id"`
	Number    int    `json:"number"`
	Fact      string `json:"fact"`
	CreatedAt string `json:"created_at"`
}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
