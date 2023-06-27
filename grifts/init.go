package grifts

import (
	"fmt"
	gqlgen "github.com/99designs/gqlgen/api"
	gqlgenConfig "github.com/99designs/gqlgen/codegen/config"
	. "github.com/gobuffalo/grift/grift"
)

var _ = Namespace("gqlgen", func() {

	var _ = Desc("generate", "Generate GraphQL types with gqlgen from the schema in graph/schema.graphqls")
	var _ = Add("generate", func(c *Context) error {
		fmt.Println("Generating GraphQL types with gqlgen...")
		cfg, err := gqlgenConfig.LoadConfig("./config/gqlgen.yml")
		if err != nil {
			return err
		}
		if err := gqlgen.Generate(cfg); err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	})

})
