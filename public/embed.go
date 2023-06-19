package dist

import (
	"embed"
	"io/fs"

	"github.com/gobuffalo/buffalo"
)

//go:embed *
var files embed.FS

// noinspection GoUnusedExportedFunction
func FS() fs.FS {
	return buffalo.NewFS(files, "dist")
}
