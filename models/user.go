package models

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

// User is a generated model from buffalo-auth, it serves as the base for username/password authentication.
type User struct {
	ID             int64     `db:"id"`      // The database ID of the user (do not expose to users!).
	NanoID         NanoID    `db:"nano_id"` // NanoID is the user-facing ID of the performance, generated using Go Nanoid.
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	EmailConfirmed bool      `db:"email_confirmed"`
	PasswordHash   string    `db:"password_hash"`

	Password string `json:"-" db:"-"` // Needed for form binding on create/update
}

func (u User) TableName() string {
	return "vp_user"
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	nanoid, err := NewNanoID()

	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.NanoID = nanoid
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.PasswordHash = string(ph)
	return tx.ValidateAndCreate(u)
}

// String is not required by pop and may be deleted
func (u User) String() string {
	return fmt.Sprintf("@%s <%s> #%s", u.Username, u.Email, u.NanoID)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
		// check to see if the username or email address is already taken:
		// we are intentionally vague about which one it is to prevent doxxing attacks
		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "credentials",
			Message: "username or email already taken%.s", // "%.s" discards any extra string added py Sprintf()
			Fn: func() bool {
				var b bool
				q := tx.Where("username = ? OR email = ?", u.Username, u.Email)
				if u.ID != 0 {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(*pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(*pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
