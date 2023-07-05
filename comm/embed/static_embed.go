package embed

import (
	"embed"
	"io/fs"
)

var staticFs embed.FS

func SetStaticFsFs(fs embed.FS) {
	staticFs = fs
}

func GetStaticFs() embed.FS {
	return staticFs
}

func GetStaticFileSystem(p ...string) fs.FS {
	path := "static"
	if len(p) != 0 && p[0] != "" {
		path = p[0]
	}
	fsys, _ := fs.Sub(GetStaticFs(), path)
	return fsys
}
