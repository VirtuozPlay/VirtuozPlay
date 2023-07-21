package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/gobuffalo/pop/v6"
)

type Resolver struct {
	DB *pop.Connection
}
