package models

import (
	"log"
	"os"

	"github.com/gobuffalo/pop"
	"github.com/joho/godotenv"
)

var tx *pop.Connection
var err error

func init() {
	godotenv.Load("../.env")
	env := os.Getenv("APP_MODE")
	var err error
	tx, err = pop.Connect(env)
	if err != nil {
		log.Panic("error connecting to database", err)
	}
}
