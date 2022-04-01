package controller

import (
	"net/http"
	"os"
	"runtime/debug"
	"time"

	log "github.com/go-kit/kit/log"

	_ "github.com/kasumusof/movies/docs" 

	"github.com/gorilla/mux"
	"github.com/kasumusof/movies/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @contact.name   MOVIES API

// router the multiplexer
var router = newMuxSever()
var a = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
var logger = log.With(a, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

// NewMuxSever creates a new mux server
func newMuxSever() *mux.Router {
	return mux.NewRouter()
}

// GetRouter dfd
func GetRouter() *mux.Router {

	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	// GET MOVIES godoc
	// @Summary		get all movies
	// @Description  get all movies
	// @Tags         movies
	// @Accept       json
	// @Produce      json
	// @Param        id   path      int  true  "Account ID"
	// @Success      200  {object}  model.Account
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /movies/ [get]
	router.HandleFunc("/movies/", getMovies).Methods("GET").Name("movies")
	router.HandleFunc("/movies/{movie_id:[0-9]+}/", getMovie).Methods("GET").Name("movie")
	router.HandleFunc("/movies/{movie_id:[0-9]+}/comments/", addComments).Methods("POST").Name("addComment")
	router.HandleFunc("/movies/{movie_id:[0-9]+}/comments/", getComments).Methods("GET").Name("comments")
	router.HandleFunc("/movies/{movie_id}/characters/", getCharacters).Methods("GET").Name("characers")
	router.HandleFunc("/characters/{character_id:[0-9]+}/", getCharacter).Methods("GET").Name("character")

	router.Use(jsonify)
	router.Use(loggingMiddleware(logger))

	return router
}

// jsonify used to set the header to application/json
func jsonify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// loggingMiddleware logs the incoming HTTP request & its duration.
func loggingMiddleware(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					utils.MakeResponse(&w, "Internal Server Error: FLM", false, http.StatusInternalServerError)

					logger.Log(
						"err", err,
						"trace", string(debug.Stack()),
						"Path", r.URL.String(),
					)
				}
			}()

			start := time.Now()
			wrapped := utils.WrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Log(
				"Status", wrapped.Status(),
				"Device", r.Header["User-Agent"],
				"Method", r.Method,
				"Path", r.URL.EscapedPath(),
				"Duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}
