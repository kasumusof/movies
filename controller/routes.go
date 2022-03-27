package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// router the multiplexer
var router = newMuxSever()

// NewMuxSever creates a new mux server
func newMuxSever() *mux.Router {
	return mux.NewRouter()
}

// GetRouter dfd
func GetRouter() *mux.Router {
	// movies := router.NewRoute().Subrouter().Name("movies").PathPrefix("/movies")
	router.HandleFunc("/movies/", getMovies).Methods("GET")
	router.HandleFunc("/movies/{movie_id:[0-9]+}", getMovie).Methods("GET")
	router.HandleFunc("/movies/{movie_id:[0-9]+}/comments/", addComments).Methods("POST")
	router.HandleFunc("/movies/{movie_id:[0-9]+}/comments/", getComments).Methods("GET")
	router.HandleFunc("/movies/{movie_id}/characters/", getCharacters).Methods("GET")
	// router.HandleFunc("/characters/{character_id:[0-9]+}", getCharacter).Methods("GET")

	router.Use(commonMiddleware)

	return router
}

// middleware used to set the header to application/json
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
