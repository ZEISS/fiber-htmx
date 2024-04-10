package htmx

import (
	"embed"
	"io/fs"
)

//go:embed static/*
var assets embed.FS

// Static returns the static assets.
func Static() fs.FS {
	f, _ := fs.Sub(assets, "static")

	return f
}
