package embed

import (
	"embed"
	"io/fs"
)

var webFs embed.FS

func SetWebFs(fs embed.FS) {
	webFs = fs
}

func GetWebFs() embed.FS {
	return webFs
}

func GetWebFileSystem(p ...string) fs.FS {
	path := "config"
	if len(p) != 0 && p[0] != "" {
		path = p[0]
	}
	fsys, _ := fs.Sub(GetWebFs(), path)
	return fsys
}
