package actions

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/ed25519"
	"net/http"
	"time"
	"virtuozplay/models/repository"

	"github.com/gobuffalo/buffalo"
	csrf "github.com/gobuffalo/mw-csrf"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"virtuozplay/models"
)

var (
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
)

func init() {
	// FIXME: this generates a new key every time the server starts.
	var err error
	publicKey, privateKey, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
}

// LogInSignupPage shows the login/signup page
func LogInSignupPage(c buffalo.Context) error {
	return csrf.New(HomeHandler)(c)
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c buffalo.Context) error {
	params := &NewUser{}
	if err := c.Bind(params); err != nil {
		return errors.WithStack(err)
	}

	users := c.Value("users").(*repository.Users)
	u, err := users.New(params.Username, params.Email, params.Password)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(models.UnwrapErrors(err)))
	}

	return respondWithJWT(c, u, http.StatusCreated)
}

type LogInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LogIn(c buffalo.Context) error {
	params := &LogInParams{}
	if err := c.Bind(params); err != nil {
		return errors.WithStack(err)
	}

	users := c.Value("users").(*repository.Users)
	u, err := users.FindByEmail(params.Email)

	// helper function to handle bad attempts
	bad := func() error {
		verrs := validate.NewErrors()
		verrs.Add("credentials", "invalid email or password")

		return c.Render(http.StatusUnauthorized, r.JSON(verrs))
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// couldn't find a user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(params.Password))
	if err != nil {
		return bad()
	}

	return respondWithJWT(c, u, http.StatusOK)
}

func LogOut(c buffalo.Context) error {
	// TODO: invalidate the JWT, see https://stackoverflow.com/questions/61368789/can-i-force-a-logout-or-expiration-of-a-jwt-token
	c.Session().Clear()
	return c.Render(http.StatusOK, r.JSON(nil))
}

const tokenExpireDuration = time.Minute * 1

func generateJWT(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"exp":        time.Now().Add(tokenExpireDuration).UnixMilli(),
			"authorized": true,
			"dbId":       u.ID,
			"nanoId":     u.NanoID,
		})
	return token.SignedString(privateKey)
}

func respondWithJWT(c buffalo.Context, u *models.User, statusCode int) error {
	token, err := generateJWT(u)

	if err != nil {
		app.Logger.Errorf("error generating JWT for user %v (%v): %v", u.Email, u.ID, err)
		return fmt.Errorf("encountered an error while logging in")
	}

	return c.Render(statusCode, r.JSON(map[string]string{
		"token":    token,
		"username": u.Username,
		"email":    u.Email,
	}))
}
