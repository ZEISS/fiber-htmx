package htmx

// HxExtType is a type for htmx extension types.
type HxExtType string

const (
	HxExtAlpineMorph         HxExtType = "alpine-morph"
	HxExtClassTools          HxExtType = "class-tools"
	HxExtClientSideTemplates HxExtType = "client-side-templates"
	HxExtIgnoreDebug         HxExtType = "ignore:debug"
	HxExtJSON                HxExtType = "json-enc"
	HxExtMultiSwap           HxExtType = "multi-swap"
	HxExtPathDeps            HxExtType = "path-deps"
)

// String returns the string representation of the htmx extension type.
func (v HxExtType) String() string {
	return string(v)
}

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

// HyperScript sets the _ attribute to specify the hyperscript for the response.
func HyperScript(v string) Node {
	return Attribute("_", v)
}
