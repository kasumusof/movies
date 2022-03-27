package main

import (
	"net/http"
	"time"

	"github.com/kasumusof/movies/controller"
)

func main() {
	r := controller.GetRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}

// Todo convert urls
// Add Redis cache
// Add a logger middleware
// Containerize
// Documentation