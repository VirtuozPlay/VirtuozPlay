package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"virtuozplay/models/repository"
)

type Resolver struct {
	Performances *repository.Performances
	Songs        *repository.Songs
	Users        *repository.Users
}
