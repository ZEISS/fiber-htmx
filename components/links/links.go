package links

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
)

// LinkProps represents the properties for a link element.
type LinkProps struct {
	Rel        string          // The relationship between the current document and the linked document.
	Href       string          // The URL of the linked document.
	ClassNames map[string]bool // The class names for the link element.

	ctx context.Context
}

// WithContext returns the LinkProps with the provided context.
func (p LinkProps) WithContext(ctx context.Context, children ...htmx.Node) LinkProps {
	p.ctx = ctx

	return p
}

// Link generates a link element based on the provided properties.
func Link(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Primary generates a primary link element based on the provided properties.
func Primary(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":         true,
			"link-primary": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Secondary generates a secondary link element based on the provided properties.
func Secondary(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":           true,
			"link-secondary": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Accent generates an accent link element based on the provided properties.
func Accent(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":        true,
			"link-accent": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Neutral generates a neutral link element based on the provided properties.
func Neutral(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":         true,
			"link-neutral": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Info generates an info link element based on the provided properties.
func Info(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":      true,
			"link-info": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Warning generates a warning link element based on the provided properties.
func Warning(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":         true,
			"link-warning": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Error generates an error link element based on the provided properties.
func Error(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":       true,
			"link-error": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}

// Underline generates an underline link element based on the provided properties.
func Underline(p LinkProps, children ...htmx.Node) htmx.Node {
	return htmx.A(
		htmx.ClassNames{
			"link":       true,
			"link-hover": true,
		}.Merge(p.ClassNames),
		htmx.Href(p.Href),
		htmx.Rel(p.Rel),
		htmx.Group(children...),
	)
}
