package models

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

type NanoID string

// NewNanoID generates a new (99.99999%) unique id.
func NewNanoID(length ...int) (NanoID, error) {
	nanoID, err := gonanoid.New(length...)

	if err != nil {
		return "", err
	}
	return NanoID(nanoID), nil
}

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}
