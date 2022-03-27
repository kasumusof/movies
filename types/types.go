package types

import "github.com/kasumusof/movies/models"

// MovieResp .
type MovieResp struct {
	Title         string          `json:"title"`
	EpisodeID     int             `json:"episode_id"`
	OpeningCrawl  string          `json:"opening_crawl"`
	Director      string          `json:"director"`
	Producer      string          `json:"producer"`
	CharacterURLs []string        `json:"characters"`
	Created       string          `json:"created"`
	Edited        string          `json:"edited"`
	URL           string          `json:"url"`
	Comments      models.Comments `json:"comment"`
	// MovieID       string          `json:"movie_id"`
}

// MoviesResp .
type MoviesResp struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []MovieResp `json:"results"`
}

// CharacterResp
type CharacterResp struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	URL       string   `json:"url"`
}

// CharactersResp
type CharactersResp struct {
	Count    int             `json:"count"`
	Next     interface{}     `json:"next"`
	Previous interface{}     `json:"previous"`
	Results  []CharacterResp `json:"results"`
}

