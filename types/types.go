package types

import "github.com/kasumusof/movies/models"

// Movie .
type Movie struct {
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

// Movies .
type Movies struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Movie     `json:"results"`
}

// Character
type Character struct {
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

// Characters .
type Characters struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Character `json:"results"`
}

// Response .
type Response struct {
	OK     bool        `json:"ok"`
	Count  int         `json:"count,omitempty"`
	Result interface{} `json:"result"`
}

// MoviesResp
type MoviesResp struct {
	OK     bool   `json:"ok"`
	Result Movies `json:"result"`
}

// MovieResp
type MovieResp struct {
	OK     bool  `json:"ok"`
	Result Movie `json:"result"`
}

// CharactersResp
type CharactersResp struct {
	OK     bool       `json:"ok"`
	Result Characters `json:"result"`
}

// CharacterResp
type CharacterResp struct {
	OK     bool       `json:"ok"`
	Result Characters `json:"result"`
}

// CommentResp
type CommentResp struct {
	OK     bool           `json:"ok"`
	Result models.Comment `json:"result"`
}

// CommentsResp
type CommentsResp struct {
	OK     bool            `json:"ok"`
	Result models.Comments `json:"result"`
}
