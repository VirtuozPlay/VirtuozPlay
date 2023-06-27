package actions

import (
	"net/http"
)

func (as *ActionSuite) Test_GraphQLHandler() {
	type virtuozPlayData struct {
		Version string `json:"version"`
	}
	type virtuozPlayContainer struct {
		VirtuozPlay virtuozPlayData `json:"virtuozPlay"`
	}
	type queryResponse[T any] struct {
		Data T `json:"data"`
	}

	res := as.JSON("/graphql").Post(map[string]interface{}{
		"query": `{
			virtuozPlay {
				version
			}
		}`,
	})

	var data queryResponse[virtuozPlayContainer]
	res.Bind(&data)

	as.Equal(http.StatusOK, res.Code)
	as.Equal("0.1.0", data.Data.VirtuozPlay.Version)
}
