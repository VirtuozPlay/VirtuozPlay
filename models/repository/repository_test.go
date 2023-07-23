package repository

import (
	"github.com/gobuffalo/suite/v4"
	"os"
	"testing"
)

type RepositorySuite struct {
	*suite.Model
	Performances Performances
}

func Test_RepositorySuite(t *testing.T) {
	model, err := suite.NewModelWithFixtures(os.DirFS("../../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &RepositorySuite{
		Model:        model,
		Performances: NewPerformancesRepository(model.DB),
	}
	suite.Run(t, as)
}
