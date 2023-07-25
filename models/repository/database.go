package repository

import (
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"virtuozplay/models"
)

// DatabaseRepository is a generic database-based repository that is used as a base for other repos like Songs and Performances.
type DatabaseRepository[V models.Value] struct {
	db *pop.Connection
}

// NewDatabaseRepository creates a NewSongsRepository instance for the given database connection.
func NewDatabaseRepository[V models.Value](db *pop.Connection) DatabaseRepository[V] {
	return DatabaseRepository[V]{db: db}
}

// FindByNanoID finds a model instance by its nanoID.
// preloadFields is a list of associations to load eagerly. (see https://gobuffalo.io/documentation/database/relations/#load-specific-associations)
// when preloadFields is empty, no associations are preloaded.
func (r *DatabaseRepository[V]) FindByNanoID(nanoID models.NanoID, preloadFields ...string) (*V, error) {
	var err error
	value := new(V)

	query := r.db.Where("nano_id = ?", nanoID)

	if p, ok := (any(value)).(models.PreLoadable); ok {
		preloadFields = p.ResolvePreloads(preloadFields...)
	} else {
		preloadFields = []string{}
	}

	if len(preloadFields) > 0 {
		query = query.EagerPreload(preloadFields...)
	}

	if err = query.First(value); err != nil {
		return nil, fmt.Errorf("%v '%v' not found", (*value).TableName(), nanoID)
	}
	return value, nil
}

// Create saves the given value to the database.
func (r *DatabaseRepository[V]) Create(value *V) error {
	return models.WrapValidation(r.db.ValidateAndCreate(value))
}

// Update updates the given value in the database.
func (r *DatabaseRepository[V]) Update(value *V) error {
	return models.WrapValidation(r.db.ValidateAndUpdate(value))
}
