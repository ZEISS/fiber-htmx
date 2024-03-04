package htmx

const (
	HxExtAlpineMorph         = "alpine-morph"
	HxExtClassTools          = "class-tools"
	HxExtClientSideTemplates = "client-side-templates"
	HxExtIgnoreDebug         = "ignore:debug"
	HxExtJSON                = "json-enc"
	HxExtMultiSwap           = "multi-swap"
	HxExtPathDeps            = "path-deps"
)

// MustacheTemplate sets the mustache-template attribute to specify the mustache template for the response.
func MustacheTemplate(v string) Node {
	return Attribute("mustache-template", v)
}

// HandlebarsTemplate sets the handlebars-template attribute to specify the handlebars template for the response.
func HandlebarsTemplate(v string) Node {
	return Attribute("handlebars-template", v)
}

// NunjucksTemplate sets the nunjucks-template attribute to specify the nunjucks template for the response.
func NunjucksTemplate(v string) Node {
	return Attribute("nunjucks-template", v)
}

// XSLTTemplat sets the xslt-template attribute to specify the xslt template for the response.
func XSLTTemplat(v string) Node {
	return Attribute("xslt-template", v)
}
