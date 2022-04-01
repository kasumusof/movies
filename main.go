package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kasumusof/movies/controller"
)

func main() {
	// @title MOVIES API documentation
	// @version 1.0.0
	// @host localhost:3000
	// @BasePath /
	port := os.Getenv("PORT")
	fmt.Println("Running")
	r := controller.GetRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("exited")
}

// Todo convert urls **
// Add Redis  **
// Add a logger middleware **
// Containerize
// Documentation
// catch errors for non-existing movies  *
// characters **
// deploy
