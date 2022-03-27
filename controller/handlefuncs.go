package controller

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kasumusof/movies/models"
	"github.com/kasumusof/movies/types"
	"github.com/kasumusof/movies/utils"
)

var (
	url          = "https://swapi.dev/api/"
	movieURL     = url + "films/"
	characterURL = url + "people/"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	url := movieURL

	var data types.MoviesResp

	x, err := utils.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(x.Body).Decode(&data)

	// load comments

	for i := 0; i < len(data.Results); i++ {
		movieURL := data.Results[i].URL
		movieID := strings.Split(movieURL, "/")[5]
		id, _ := strconv.Atoi(movieID)
		comments, err := models.GetComments(id)
		if err != nil {
			// do the error response
		}
		data.Results[i].Comments = *comments
	}

	// load comments

	resp, err := json.Marshal(data)
	w.Write(resp)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]
	url := movieURL + movieID

	var data types.MovieResp

	x, err := utils.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(x.Body).Decode(&data)

	// load comments
	id, _ := strconv.Atoi(movieID)
	comments, err := models.GetComments(id)
	if err != nil {
		// do the error response
	}
	data.Comments = *comments

	resp, err := json.Marshal(data)
	w.Write(resp)
}

func getCharacters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	url := movieURL + movieID
	var data types.MovieResp

	x, err := utils.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(x.Body).Decode(&data)

	type characters []types.CharacterResp

	chars := characters{}

	for i := 0; i < len(data.CharacterURLs); i++ {
		char := types.CharacterResp{}
		req, err := utils.Fetch(data.CharacterURLs[i])
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(req.Body).Decode(&char)

		chars = append(chars, char)
	}

	resp, err := json.Marshal(chars)
	w.Write(resp)
}

func getCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterID := vars["character_id"]
	url := characterURL + characterID

	var data types.CharacterResp

	x, err := utils.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(x.Body).Decode(&data)

	resp, err := json.Marshal(data)
	w.Write(resp)
}

func addComments(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	movieID := mux.Vars(r)["movie_id"]
	r.ParseForm()
	text := r.Form.Get("text")

	id, _ := strconv.Atoi(movieID)
	comment.MovieID = id
	ip,_,_ := net.SplitHostPort(r.RemoteAddr)
	comment.Author = ip
	comment.Text = text
	log.Println(ip)

	a, _ := models.NewComment(comment)

	resp, _ := json.Marshal(a)
	w.Write(resp)

}

func getComments(w http.ResponseWriter, r *http.Request) {
	var comments *models.Comments
	movieID := mux.Vars(r)["movie_id"]
	r.ParseForm()

	id, _ := strconv.Atoi(movieID)
	
	comments, _ = models.GetComments(id)

	resp, _ := json.Marshal(comments)
	w.Write(resp)
}
