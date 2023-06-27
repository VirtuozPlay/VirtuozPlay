package actions

import (
	"io/fs"
	"virtuozplay/dist"
)

func DistFS() fs.FS {
	return dist.FS()
}
