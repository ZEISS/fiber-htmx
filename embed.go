package htmx

import (
	"embed"
)

const (
	BundleFolder      = "dist"
	DefaultBundleFile = "out.js"
)

//go:embed dist
var Bundle embed.FS
