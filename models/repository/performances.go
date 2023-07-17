package repository

import (
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"time"
	"virtuozplay/models"
)

// Performances is the repository for the Performance model.
// Provides methods for find and CRUD operations of performances.
type Performances struct {
	db         *pop.Connection
	inProgress map[models.NanoID]*models.Performance
}

func NewPerformancesRepository(db *pop.Connection) Performances {
	return Performances{db: db, inProgress: make(map[models.NanoID]*models.Performance)}
}

// FindByNanoID finds a performance by its nanoID.
func (r *Performances) FindByNanoID(nanoID models.NanoID) (*models.Performance, error) {
	if perf, isInProgress := r.inProgress[nanoID]; isInProgress {
		return perf, nil
	}

	var err error
	perf := &models.Performance{}

	query := r.db.Where("nano_id = ?", nanoID)
	if err = query.First(perf); err != nil {
		return nil, fmt.Errorf("performance %v not found", nanoID)
	}
	return perf, nil
}

// FindInProgressByNanoID finds a performance by its nanoID only if it stored in memory (a.K.a. in progress)
func (r *Performances) FindInProgressByNanoID(nanoID models.NanoID) (*models.Performance, error) {
	perf, isInProgress := r.inProgress[nanoID]

	if !isInProgress {
		return nil, fmt.Errorf("performance %v does not exist or is not in progress", nanoID)
	}
	return perf, nil
}

// Create saves the given performance to the database.
func (r *Performances) Create(perf *models.Performance) *models.ValidationErrors {
	return models.WrapValidation(r.db.ValidateAndCreate(perf))
}

// New creates a new in-progress performance with a unique nanoID (without saving it to the database).
func (r *Performances) New(nanoIDLen ...int) (*models.Performance, error) {
	nanoId, err := models.NewNanoID(nanoIDLen...)

	if err != nil {
		return nil, fmt.Errorf("failed to generate unique id for performance")
	}
	perf := &models.Performance{NanoID: nanoId}
	perf.CreatedAt = time.Now().Truncate(time.Microsecond)

	r.inProgress[nanoId] = perf
	return perf, nil
}

// Update updates the given performance in the database.
func (r *Performances) Update(perf *models.Performance) *models.ValidationErrors {
	if _, isInProgress := r.inProgress[perf.NanoID]; isInProgress {
		return models.WrapValidation(perf.Validate(nil))
	}
	return models.WrapValidation(r.db.ValidateAndUpdate(perf))
}

// MarkAsFinished marks the given performance as finished by removing it from the in-progress map and saving it to the database.
func (r *Performances) MarkAsFinished(perf *models.Performance) *models.ValidationErrors {
	if _, isInProgress := r.inProgress[perf.NanoID]; !isInProgress {
		return models.WrapValidation(nil, fmt.Errorf("performance %v does not exist or is not in progress", perf.NanoID))
	}
	if err := models.WrapValidation(r.db.ValidateAndSave(perf)); err != nil {
		return err
	}
	delete(r.inProgress, perf.NanoID)
	return nil
}