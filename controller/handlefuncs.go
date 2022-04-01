package controller

import (
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"

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

// GET MOVIES godoc
// @Summary		get all movies
// @Description  get all movies
// @Tags         movies
// @Produce      json
// @Success      200  {object}  types.MoviesResp
// @Router       /movies/ [get]
func getMovies(w http.ResponseWriter, r *http.Request) {
	url := movieURL

	var data types.Movies
	// utils.FetchCachewithComments()
	utils.CheckAndCache(url, 0, &data)

	// check my own cache
	long, err := utils.CheckMyCache(r.URL.String(), 0, data)
	if err == nil {
		// means my cache exists
		data = long.(types.Movies)
	} else {
		// replace urls for characters and url
		rout := GetRouter()
		re := regexp.MustCompile("[0-9]+")
		for i := 0; i < len(data.Results); i++ {
			for j := 0; j < len(data.Results[i].CharacterURLs); j++ {
				charID := re.FindAllString(data.Results[i].CharacterURLs[j], -1)
				uRl, err := rout.Get("character").URL("character_id", charID[0])
				if err != nil {
					log.Fatal(err)
				}
				data.Results[i].CharacterURLs[j] = uRl.String()
			}
			id := re.FindAllString(data.Results[i].URL, -1)[0]
			x, _ := rout.Get("movie").URL("movie_id", id)
			data.Results[i].URL = x.String()
		}
		utils.MakeMyCache(r.URL.String(), 0, data)
	}

	// load comments

	for i := 0; i < len(data.Results); i++ {
		movieURL := data.Results[i].URL
		re := regexp.MustCompile("[0-9]+")
		movieID := re.FindAllString(movieURL, -1)[0]
		id, _ := strconv.Atoi(movieID)
		comments, err := models.GetComments(id)
		if err != nil {
			// do the error response
		}
		data.Results[i].Comments = *comments
	}

	// load comments end

	utils.MakeResponse(&w, data, true, http.StatusOK)
}

// GET MOVIES godoc
// @Summary		get a particular movies
// @Description  get a particular movies
// @Tags         movies
// @Param        movie_id   path      int  true "Movie ID"
// @Produce      json
// @Success      200  {object}  types.MovieResp
// @Router       /movies/{movie_id}/ [get]
func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]
	url := movieURL + movieID
	var data types.Movie

	// check cache else fetch
	utils.CheckAndCache(url, 0, &data)

	// just to check error codes
	x, err := utils.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	if x.StatusCode != http.StatusOK {
		utils.MakeResponse(&w, "Not Found", false, x.StatusCode)
		return
	}

	// check my own cache
	long, err := utils.CheckMyCache(r.URL.String(), 0, data)
	if err == nil {
		// means my cache exists
		data = long.(types.Movie)

	} else {
		// replace urls for characters and url
		rout := GetRouter()
		for i := 0; i < len(data.CharacterURLs); i++ {
			re := regexp.MustCompile("[0-9]+")
			charID := re.FindAllString(data.CharacterURLs[i], -1)
			uRl, err := rout.Get("character").URL("character_id", charID[0])
			if err != nil {
				log.Fatal(err)
			}
			data.CharacterURLs[i] = uRl.String()
		}

		y, _ := rout.Get("movie").URL("movie_id", movieID)
		data.URL = y.String()
		utils.MakeMyCache(r.URL.String(), 0, data)
	}

	// load comments
	id, _ := strconv.Atoi(movieID)
	comments, err := models.GetComments(id)
	if err != nil {
		log.Println(err)
	}
	data.Comments = *comments

	utils.MakeResponse(&w, data, true, http.StatusOK)
}

// GET MOVIES godoc
// @Summary		get all  characters for a particular movie
// @Description  get all  characters for a particular movie
// @Tags         movies
// @Param        movie_id   path      int  true "Movie ID"
// @Param        character_id   path      int  true "Character ID"
// @Produce      json
// @Success      200  {object}  types.CharactersResp
// @Router       /movies/{movie_id}/characters/{character_id}/ [get]
func getCharacters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	url := movieURL + movieID
	var data types.Movie
	// r.ParseForm()
	// gender := r.Form.Get("gender")

	utils.CheckAndCache(url, 0, &data)

	// check for errors from fetch (not the best implementation)
	x, err := utils.Fetch(url)
	if err != nil {
		utils.MakeResponse(&w, "error", true, http.StatusInternalServerError)
		return
	}
	if x.StatusCode != http.StatusOK {
		utils.MakeResponse(&w, "Not found", false, x.StatusCode)
		return
	}

	type characters []types.Character

	chars := characters{}

	for i := 0; i < len(data.CharacterURLs); i++ {
		char := types.Character{}
		utils.CheckAndCache(data.CharacterURLs[i], 0, &char)

		// replace urls for movies and url
		rout := GetRouter()
		re := regexp.MustCompile("[0-9]+")
		for i := 0; i < len(char.Films); i++ {
			id := re.FindAllString(char.Films[i], -1)[0]
			uRl, err := rout.Get("movie").URL("movie_id", id)
			if err != nil {
				log.Println(err)
			}
			char.Films[i] = uRl.String()
		}
		id := re.FindAllString(char.URL, -1)[0]
		y, _ := rout.Get("character").URL("character_id", id)
		data.URL = y.String()
		// end of urls
		chars = append(chars, char)
	}

	utils.MakeResponse(&w, chars, true, http.StatusOK)
}

// GET MOVIES godoc
// @Summary		get a particular character
// @Description  get a particular character
// @Tags         movies
// @Param        character_id   path      int  true "Character ID"
// @Produce      json
// @Success      200  {object}  types.CharacterResp
// @Router       /character/{character_id}/ [get]
func getCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterID := vars["character_id"]
	url := characterURL + characterID

	var data types.Character

	utils.CheckAndCache(url, 0, &data)

	x, err := utils.Fetch(url)
	if err != nil {
		utils.MakeResponse(&w, "error", false, http.StatusInternalServerError)
		return
	}

	if x.StatusCode != http.StatusOK {
		utils.MakeResponse(&w, "Not Found", false, x.StatusCode)
		return
	}

	// replace urls for movies and url
	rout := GetRouter()
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(data.Films); i++ {
		id := re.FindAllString(data.Films[i], -1)[0]
		uRl, err := rout.Get("movie").URL("movie_id", id)
		if err != nil {
			log.Println(err)
		}
		data.Films[i] = uRl.String()
	}

	y, _ := rout.Get("character").URL("character_id", characterID)
	data.URL = y.String()
	// end of urls

	utils.MakeResponse(&w, data, true, http.StatusOK)
}

// GET MOVIES godoc
// @Summary		add a comment to a movie
// @Description  add a comment to a movie
// @Tags         movies
// @Param        movie_id   path      int  true "Movie ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.CommentResp
// @Router       /movies/{movie_id}/comments/ [post]
func addComments(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	movieID := mux.Vars(r)["movie_id"]
	r.ParseForm()
	text := r.Form.Get("text")
	if len(text) > 500 {
		utils.MakeResponse(&w, "Maximum of 500 characters per comment", false, http.StatusNotAcceptable)
		return
	}
	id, _ := strconv.Atoi(movieID)
	comment.MovieID = id
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	comment.Author = ip
	comment.Text = text
	log.Println(ip)

	a, err := models.NewComment(comment)
	if err != nil {
		log.Println(err)
	}

	utils.MakeResponse(&w, a, true, http.StatusOK)
}

// GET MOVIES godoc
// @Summary		get all comments for  a movie
// @Description  get all comments for a movie
// @Tags         movies
// @Param        movie_id   path      int  true "Movie ID"
// @Produce      json
// @Success      200  {object}  types.CommentsResp
// @Router       /movies/{movie_id}/comments/ [get]
func getComments(w http.ResponseWriter, r *http.Request) {
	var comments *models.Comments
	movieID := mux.Vars(r)["movie_id"]
	r.ParseForm()

	id, _ := strconv.Atoi(movieID)

	comments, err := models.GetComments(id)
	if err != nil {
		log.Println(err)
	}

	utils.MakeResponse(&w, comments, true, http.StatusOK)
}
