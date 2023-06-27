package actions

import (
	"github.com/gobuffalo/httptest"
	"os"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

// JSON shadows suite.Action.JSON to set the Content-Type and Accept headers to application/json.
func (as *ActionSuite) JSON(u string, args ...interface{}) *httptest.JSON {
	h := httptest.New(as.App)

	h.Headers["Content-Type"] = "application/json"
	h.Headers["Accept"] = "application/json"
	return h.JSON(u, args...)
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), os.DirFS("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
