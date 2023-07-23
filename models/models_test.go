package models

import (
	"os"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ModelSuite struct {
	*suite.Model
}

func (ms *ModelSuite) Test_NewNanoID() {
	nanoID, err := NewNanoID()

	ms.NoError(err)
	ms.NotEmpty(nanoID)
	ms.Equal(21, len(nanoID))

	nanoID, err = NewNanoID(-1)
	ms.Empty(nanoID)
	ms.Error(err)
}

func Test_ModelSuite(t *testing.T) {
	model, err := suite.NewModelWithFixtures(os.DirFS("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ModelSuite{
		Model: model,
	}
	suite.Run(t, as)
}
