package repository

import (
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"sync"
	"time"
	"virtuozplay/models"
)

// Performances is the repository for the Performance model.
// Provides methods for find and CRUD operations of performances.
type Performances struct {
	DatabaseRepository[models.Performance]
	inProgress sync.Map
}

func NewPerformancesRepository(db *pop.Connection) Performances {
	return Performances{
		DatabaseRepository: NewDatabaseRepository[models.Performance](db),
		inProgress:         sync.Map{},
	}
}

// FindByNanoID finds a performance by its nanoID.
// preloadFields is a list of associations to load eagerly. (see https://gobuffalo.io/documentation/database/relations/#load-specific-associations)
// when preloadFields is empty, no associations are preloaded.
func (r *Performances) FindByNanoID(nanoID models.NanoID, preloadedFields ...string) (*models.Performance, error) {
	if perf, isInProgress := r.inProgress.Load(nanoID); isInProgress {
		return perf.(*models.Performance), nil
	}

	return r.DatabaseRepository.FindByNanoID(nanoID, preloadedFields...)
}

// FindInProgressByNanoID finds a performance by its nanoID only if it stored in memory (a.k.a. in progress)
func (r *Performances) FindInProgressByNanoID(nanoID models.NanoID) (*models.Performance, error) {
	perf, isInProgress := r.inProgress.Load(nanoID)

	if !isInProgress {
		return nil, fmt.Errorf("performance %v does not exist or is not in progress", nanoID)
	}
	return perf.(*models.Performance), nil
}

// New creates a new in-progress performance with a unique nanoID (without saving it to the database).
func (r *Performances) New(song *models.Song, nanoIDLen ...int) (*models.Performance, error) {
	nanoId, err := models.NewNanoID(nanoIDLen...)

	if err != nil {
		return nil, fmt.Errorf("failed to generate unique id for performance")
	}
	perf := &models.Performance{
		NanoID:    nanoId,
		SongID:    song.ID,
		Song:      song,
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}

	r.inProgress.Store(nanoId, perf)
	return perf, nil
}

// Update updates the given performance in the database.
func (r *Performances) Update(perf *models.Performance) error {
	if _, isInProgress := r.inProgress.Load(perf.NanoID); isInProgress {
		return models.WrapValidation(perf.Validate(nil))
	}
	return r.DatabaseRepository.Update(perf)
}

// MarkAsFinished marks the given performance as finished by removing it from the in-progress map and saving it to the database.
func (r *Performances) MarkAsFinished(perf *models.Performance) error {
	if _, isInProgress := r.inProgress.Load(perf.NanoID); !isInProgress {
		return models.WrapValidation(nil, fmt.Errorf("performance %v does not exist or is not in progress", perf.NanoID))
	}
	if err := models.WrapValidation(r.db.ValidateAndSave(perf)); err != nil {
		return err
	}
	r.inProgress.Delete(perf.NanoID)
	return nil
}
