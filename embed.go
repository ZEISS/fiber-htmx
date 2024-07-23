package htmx

import (
	"embed"
)

const (
	BundleFoler       = "dist"
	DefaultBundleFile = "out.js"
)

//go:embed dist
var Bundle embed.FS
