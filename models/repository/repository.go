package repository

import (
	"virtuozplay/models"
)

// Repository is the constraint that all VirtuozPlay model repositories must implement.
type Repository[V any] interface {
	FindByNanoID(nanoID models.NanoID, preloadedFields ...string) (*V, error)
	Create(value *V) error
	Update(value *V) error
}
