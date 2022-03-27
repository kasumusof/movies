package models

import (
	"log"

	"github.com/gobuffalo/pop"
)

var tx *pop.Connection
var err error

func init() {
	var err error
	tx, err = pop.Connect("development")
	if err != nil {
		log.Panic("error connecting to database", err)
	}
}
