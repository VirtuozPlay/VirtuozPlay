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

// NanoID is the unique id format used throughout the application.
// Use NewNanoID to generate a new NanoID instance.
type NanoID string

// Value is the interface that all models must implement.
type Value interface {
	pop.TableNameAble
}

// The PreLoadable interface.
// When a model implements PreLoadable, the ResolvePreloads method is called to converts the `preloads` arrays from
// the GraphQL query into a list of database associations to load eagerly.
type PreLoadable interface {
	ResolvePreloads(preloads ...string) []string
}

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
