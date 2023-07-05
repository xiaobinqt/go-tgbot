package embed

import (
	"embed"
	"io/fs"
)

var configFs embed.FS

func SetConfigFs(fs embed.FS) {
	configFs = fs
}

func GetWebFs() embed.FS {
	return configFs
}

func GetConfigFileSystem(p ...string) fs.FS {
	path := "config"
	if len(p) != 0 && p[0] != "" {
		path = p[0]
	}
	fsys, _ := fs.Sub(GetWebFs(), path)
	return fsys
}
