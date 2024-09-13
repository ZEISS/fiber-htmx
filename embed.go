package htmx

import (
	"embed"
)

const (
	BundleFolder      = "dist"
	DefaultBundleFile = "fiber-htmx.min.js"
)

//go:embed dist
var Bundle embed.FS
