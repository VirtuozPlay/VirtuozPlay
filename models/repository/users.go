package repository

import (
	"github.com/gobuffalo/pop/v6"
	"virtuozplay/models"
)

// Users is the repository for the models.User instances.
// Provides methods for find and CRUD operations of performances.
type Users struct {
	DatabaseRepository[models.User]
}

// NewUsersRepository creates a new repository for users.
func NewUsersRepository(db *pop.Connection) Users {
	return Users{NewDatabaseRepository[models.User](db)}
}

func (r *Users) New(username string, email string, password string) (*models.User, error) {
	u := &models.User{
		Username:       username,
		Email:          email,
		EmailConfirmed: false,
		Password:       password,
	}

	if err := models.WrapValidation(u.Create(r.db)); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Users) FindByNanoID(nanoID models.NanoID, preloadFields ...string) (*models.User, error) {
	return r.DatabaseRepository.FindByNanoID(nanoID, preloadFields...)
}

func (r *Users) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.db.Where("email = ?", email).First(u); err != nil {
		return nil, err
	}
	return u, nil
}
