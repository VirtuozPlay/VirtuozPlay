package actions

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/suite/v4"
	"log"
	"os"
	"testing"
	"virtuozplay/models"
	"virtuozplay/models/repository"
)

type ActionSuite struct {
	*suite.Action
	Users *repository.Users
}

// JSON shadows suite.Action.JSON to set the Content-Type and Accept headers to application/json.
func (as *ActionSuite) JSON(u string, args ...interface{}) *httptest.JSON {
	h := httptest.New(as.App)

	h.Headers["Content-Type"] = "application/json"
	h.Headers["Accept"] = "application/json"
	return h.JSON(u, args...)
}

func Test_ActionSuite(t *testing.T) {
	c, err := pop.Connect(envy.Get("GO_ENV", "test"))
	if err != nil {
		log.Fatalf("could not connect to test database %v", err)
	}
	models.DB = c
	action, err := suite.NewActionWithFixtures(App(), os.DirFS("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	users := repository.NewUsersRepository(c)

	as := &ActionSuite{
		Action: action,
		Users:  &users,
	}
	suite.Run(t, as)
}
